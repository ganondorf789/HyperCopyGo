// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package trade

import (
	"context"

	"demo/api/trade/v1"
)

type ITradeV1 interface {
	TradePlaceOrder(ctx context.Context, req *v1.TradePlaceOrderReq) (res *v1.TradePlaceOrderRes, err error)
	TradeMarketClose(ctx context.Context, req *v1.TradeMarketCloseReq) (res *v1.TradeMarketCloseRes, err error)
	TradeLimitClose(ctx context.Context, req *v1.TradeLimitCloseReq) (res *v1.TradeLimitCloseRes, err error)
	TradeSetTpSl(ctx context.Context, req *v1.TradeSetTpSlReq) (res *v1.TradeSetTpSlRes, err error)
	TradeCancelOrder(ctx context.Context, req *v1.TradeCancelOrderReq) (res *v1.TradeCancelOrderRes, err error)
}
