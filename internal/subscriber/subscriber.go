package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/websocket"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	channelNewPositions = "new_positions"
	channelMarketAlert  = "market_alert"
)

func Start(ctx context.Context) {
	go runWithReconnect(ctx)
}

func runWithReconnect(ctx context.Context) {
	for {
		if err := subscribe(ctx); err != nil {
			g.Log().Errorf(ctx, "[subscriber] error: %v, reconnecting in 5s", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func subscribe(ctx context.Context) error {
	conn, err := g.Redis().Conn(ctx)
	if err != nil {
		return fmt.Errorf("redis conn: %w", err)
	}
	defer conn.Close(ctx)

	_, err = conn.Subscribe(ctx, channelNewPositions, channelMarketAlert)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	g.Log().Infof(ctx, "[subscriber] listening on channels: %s, %s", channelNewPositions, channelMarketAlert)

	for {
		msg, err := conn.ReceiveMessage(ctx)
		if err != nil {
			return fmt.Errorf("receive: %w", err)
		}

		switch msg.Channel {
		case channelNewPositions:
			handleNewPosition(ctx, msg.Payload)
		case channelMarketAlert:
			handleMarketAlert(ctx, msg.Payload)
		}
	}
}

func handleNewPosition(ctx context.Context, payload string) {
	var evt model.NewPositionEvent
	if err := json.Unmarshal([]byte(payload), &evt); err != nil {
		g.Log().Errorf(ctx, "[subscriber] parse new_position error: %v", err)
		return
	}

	id, err := dao.TraderAssetPositions.Ctx(ctx).Data(gconv.Map(evt)).InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "[subscriber] insert new_position error: %v", err)
		return
	}

	hub := websocket.GetHub()
	hub.Broadcast(websocket.WsMessage{
		Type: "new_position",
		Data: model.NewPositionBroadcast{
			Id:             id,
			Address:        evt.Address,
			Coin:           evt.Coin,
			Szi:            evt.Szi,
			LeverageType:   evt.LeverageType,
			Leverage:       evt.Leverage,
			EntryPx:        evt.EntryPx,
			PositionValue:  evt.PositionValue,
			UnrealizedPnl:  evt.UnrealizedPnl,
			ReturnOnEquity: evt.ReturnOnEquity,
			LiquidationPx:  evt.LiquidationPx,
			MarginUsed:     evt.MarginUsed,
			MaxLeverage:    evt.MaxLeverage,
		},
	})

	g.Log().Infof(ctx, "[subscriber] new_position saved and broadcast: id=%d %s %s", id, evt.Address, evt.Coin)
}

func handleMarketAlert(ctx context.Context, payload string) {
	var alert model.MarketAlert
	if err := json.Unmarshal([]byte(payload), &alert); err != nil {
		g.Log().Errorf(ctx, "[subscriber] parse market_alert error: %v", err)
		return
	}

	title := "Market Alert"
	content := fmt.Sprintf("%d new positions opened in the last %d minutes (threshold: %d)",
		alert.Count, alert.Minutes, alert.Threshold)

	id, err := dao.Notification.Ctx(ctx).Data(g.Map{
		"user_id":  0,
		"category": "market",
		"title":    title,
		"content":  content,
		"level":    1,
		"status":   1,
	}).InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "[subscriber] insert notification error: %v", err)
		return
	}

	hub := websocket.GetHub()
	hub.Broadcast(websocket.WsMessage{
		Type: "notification",
		Data: model.MarketAlertBroadcast{
			Id:       id,
			Category: "market",
			Title:    title,
			Content:  content,
			Level:    1,
		},
	})

	g.Log().Warningf(ctx, "[subscriber] MARKET ALERT broadcast: %d positions in %d min", alert.Count, alert.Minutes)
}
