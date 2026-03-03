package ws

import (
	"context"

	v1 "demo/api/ws/v1"
	wsHub "demo/internal/websocket"
	"demo/utility"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) WsConnect(ctx context.Context, req *v1.WsConnectReq) (res *v1.WsConnectRes, err error) {
	r := g.RequestFromCtx(ctx)

	claims, err := utility.ParseToken(req.Token)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 401, "message": "token无效"})
		return nil, nil
	}

	userId := claims.UserId

	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	hub := wsHub.GetHub()
	client := hub.Register(userId, ws)
	defer hub.Unregister(userId, client)

	_ = client.Write(wsHub.WsMessage{
		Type: "connected",
		Data: g.Map{"userId": userId},
	})

	g.Log().Infof(ctx, "WebSocket connected: userId=%d, online=%d", userId, hub.OnlineUserCount())

	for {
		_, msgByte, readErr := ws.ReadMessage()
		if readErr != nil {
			g.Log().Infof(ctx, "WebSocket disconnected: userId=%d", userId)
			return nil, nil
		}

		var msg wsHub.WsMessage
		if decodeErr := gjson.DecodeTo(msgByte, &msg); decodeErr != nil {
			_ = client.Write(wsHub.WsMessage{
				Type: "error",
				Data: "invalid message format",
			})
			continue
		}

		switch msg.Type {
		case "ping":
			_ = client.Write(wsHub.WsMessage{
				Type: "pong",
				Data: nil,
			})
		}
	}
}
