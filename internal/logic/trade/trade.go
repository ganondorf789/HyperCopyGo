package trade

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	hyperliquid "github.com/sonirico/go-hyperliquid"

	v1 "demo/api/trade/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterTrade(&sTrade{})
}

type sTrade struct{}

func (s *sTrade) getWallet(ctx context.Context, userId, walletId int64) (*entity.Wallet, error) {
	var w entity.Wallet
	err := dao.Wallet.Ctx(ctx).
		Where("id = ? AND user_id = ?", walletId, userId).
		Scan(&w)
	if err != nil {
		return nil, err
	}
	if w.Id == 0 {
		return nil, fmt.Errorf("钱包不存在")
	}
	if w.Status != 1 {
		return nil, fmt.Errorf("钱包已禁用")
	}
	return &w, nil
}

func (s *sTrade) newExchange(ctx context.Context, w *entity.Wallet) (*hyperliquid.Exchange, error) {
	privateKey, err := crypto.HexToECDSA(w.ApiSecretKey)
	if err != nil {
		return nil, fmt.Errorf("钱包私钥无效: %v", err)
	}
	exchange := hyperliquid.NewExchange(
		ctx,
		privateKey,
		hyperliquid.MainnetAPIURL,
		nil,
		"",
		w.Address,
		nil,
		nil,
	)
	return exchange, nil
}

func toOrderResult(s hyperliquid.OrderStatus) *v1.OrderResult {
	r := &v1.OrderResult{}
	if s.Resting != nil {
		r.Resting = &v1.OrderResultResting{
			Oid:    s.Resting.Oid,
			Status: s.Resting.Status,
		}
	}
	if s.Filled != nil {
		r.Filled = &v1.OrderResultFilled{
			TotalSz: s.Filled.TotalSz,
			AvgPx:   s.Filled.AvgPx,
			Oid:     s.Filled.Oid,
		}
	}
	if s.Error != nil {
		r.Error = s.Error
	}
	return r
}

func (s *sTrade) PlaceOrder(ctx context.Context, userId int64, in v1.TradePlaceOrderReq) (res *v1.TradePlaceOrderRes, err error) {
	w, err := s.getWallet(ctx, userId, in.WalletId)
	if err != nil {
		return nil, err
	}
	exchange, err := s.newExchange(ctx, w)
	if err != nil {
		return nil, err
	}

	req := hyperliquid.CreateOrderRequest{
		Coin:  in.Coin,
		IsBuy: in.IsBuy,
		Price: in.Price,
		Size:  in.Size,
	}

	switch in.OrderType {
	case "limit":
		req.OrderType = hyperliquid.OrderType{
			Limit: &hyperliquid.LimitOrderType{Tif: hyperliquid.Tif(in.Tif)},
		}
	case "market":
		req.OrderType = hyperliquid.OrderType{
			Limit: &hyperliquid.LimitOrderType{Tif: hyperliquid.TifIoc},
		}
	default:
		return nil, fmt.Errorf("不支持的订单类型: %s", in.OrderType)
	}

	status, err := exchange.Order(ctx, req, nil)
	if err != nil {
		return nil, fmt.Errorf("下单失败: %v", err)
	}

	return &v1.TradePlaceOrderRes{Status: toOrderResult(status)}, nil
}

func (s *sTrade) MarketClose(ctx context.Context, userId int64, in v1.TradeMarketCloseReq) (res *v1.TradeMarketCloseRes, err error) {
	w, err := s.getWallet(ctx, userId, in.WalletId)
	if err != nil {
		return nil, err
	}
	exchange, err := s.newExchange(ctx, w)
	if err != nil {
		return nil, err
	}

	var sz *float64
	if in.Size > 0 {
		sz = &in.Size
	}

	status, err := exchange.MarketClose(ctx, in.Coin, sz, nil, in.Slippage, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("市价平仓失败: %v", err)
	}

	return &v1.TradeMarketCloseRes{Status: toOrderResult(status)}, nil
}

func (s *sTrade) LimitClose(ctx context.Context, userId int64, in v1.TradeLimitCloseReq) (res *v1.TradeLimitCloseRes, err error) {
	w, err := s.getWallet(ctx, userId, in.WalletId)
	if err != nil {
		return nil, err
	}
	exchange, err := s.newExchange(ctx, w)
	if err != nil {
		return nil, err
	}

	req := hyperliquid.CreateOrderRequest{
		Coin:       in.Coin,
		IsBuy:      false,
		Price:      in.Price,
		Size:       in.Size,
		ReduceOnly: true,
		OrderType: hyperliquid.OrderType{
			Limit: &hyperliquid.LimitOrderType{Tif: hyperliquid.Tif(in.Tif)},
		},
	}

	status, err := exchange.Order(ctx, req, nil)
	if err != nil {
		return nil, fmt.Errorf("限价平仓失败: %v", err)
	}

	return &v1.TradeLimitCloseRes{Status: toOrderResult(status)}, nil
}

func (s *sTrade) SetTpSl(ctx context.Context, userId int64, in v1.TradeSetTpSlReq) (res *v1.TradeSetTpSlRes, err error) {
	w, err := s.getWallet(ctx, userId, in.WalletId)
	if err != nil {
		return nil, err
	}
	exchange, err := s.newExchange(ctx, w)
	if err != nil {
		return nil, err
	}

	req := hyperliquid.CreateOrderRequest{
		Coin:       in.Coin,
		IsBuy:      in.IsBuy,
		Price:      in.TriggerPx,
		Size:       in.Size,
		ReduceOnly: true,
		OrderType: hyperliquid.OrderType{
			Trigger: &hyperliquid.TriggerOrderType{
				TriggerPx: in.TriggerPx,
				IsMarket:  true,
				Tpsl:      hyperliquid.Tpsl(in.TpslType),
			},
		},
	}

	status, err := exchange.Order(ctx, req, nil)
	if err != nil {
		return nil, fmt.Errorf("设置止盈止损失败: %v", err)
	}

	return &v1.TradeSetTpSlRes{Status: toOrderResult(status)}, nil
}

func (s *sTrade) OpenOrders(ctx context.Context, userId int64, in v1.TradeOpenOrdersReq) (res *v1.TradeOpenOrdersRes, err error) {
	w, err := s.getWallet(ctx, userId, in.WalletId)
	if err != nil {
		return nil, err
	}

	info := hyperliquid.NewInfo(ctx, hyperliquid.MainnetAPIURL, true, nil, nil, nil)
	orders, err := info.FrontendOpenOrders(ctx, w.Address)
	if err != nil {
		return nil, fmt.Errorf("获取挂单失败: %v", err)
	}

	list := make([]v1.OpenOrderItem, 0, len(orders))
	for _, o := range orders {
		list = append(list, v1.OpenOrderItem{
			Coin:             o.Coin,
			Oid:              o.Oid,
			Side:             string(o.Side),
			LimitPx:          o.LimitPx,
			Size:             o.Sz,
			OrigSize:         o.OrigSz,
			OrderType:        o.OrderType,
			ReduceOnly:       o.ReduceOnly,
			IsTrigger:        o.IsTrigger,
			IsPositionTpSl:   o.IsPositionTpSl,
			TriggerPx:        o.TriggerPx,
			TriggerCondition: o.TriggerCondition,
			Timestamp:        o.Timestamp,
		})
	}

	return &v1.TradeOpenOrdersRes{List: list}, nil
}

func (s *sTrade) CancelOrder(ctx context.Context, userId int64, in v1.TradeCancelOrderReq) error {
	w, err := s.getWallet(ctx, userId, in.WalletId)
	if err != nil {
		return err
	}
	exchange, err := s.newExchange(ctx, w)
	if err != nil {
		return err
	}

	resp, err := exchange.Cancel(ctx, in.Coin, in.OrderId)
	if err != nil {
		return fmt.Errorf("取消订单失败: %v", err)
	}
	if !resp.Ok {
		return fmt.Errorf("取消订单失败: %s", resp.Err)
	}
	return nil
}
