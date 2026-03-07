// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/trade/v1"
)

type (
	ITrade interface {
		PlaceOrder(ctx context.Context, userId int64, in v1.TradePlaceOrderReq) (res *v1.TradePlaceOrderRes, err error)
		MarketClose(ctx context.Context, userId int64, in v1.TradeMarketCloseReq) (res *v1.TradeMarketCloseRes, err error)
		LimitClose(ctx context.Context, userId int64, in v1.TradeLimitCloseReq) (res *v1.TradeLimitCloseRes, err error)
		SetTpSl(ctx context.Context, userId int64, in v1.TradeSetTpSlReq) (res *v1.TradeSetTpSlRes, err error)
		CancelOrder(ctx context.Context, userId int64, in v1.TradeCancelOrderReq) error
	}
)

var (
	localTrade ITrade
)

func Trade() ITrade {
	if localTrade == nil {
		panic("implement not found for interface ITrade, forgot register?")
	}
	return localTrade
}

func RegisterTrade(i ITrade) {
	localTrade = i
}
