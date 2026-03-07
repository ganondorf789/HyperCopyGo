package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 下单 ====================

type TradePlaceOrderReq struct {
	g.Meta    `path:"/trade/order" tags:"Trade" method:"post" summary:"下单" login_required:"true"`
	WalletId  int64   `json:"walletId" v:"required#请选择钱包"`
	Coin      string  `json:"coin" v:"required#请输入币种"`
	IsBuy     bool    `json:"isBuy"`
	Price     float64 `json:"price" v:"required#请输入价格"`
	Size      float64 `json:"size" v:"required#请输入数量"`
	OrderType string  `json:"orderType" v:"required|in:limit,market#请选择订单类型|订单类型只能是limit或market"` // limit / market
	Tif       string  `json:"tif" d:"Gtc"`                                                        // Gtc / Alo / Ioc
}
type TradePlaceOrderRes struct {
	g.Meta `mime:"application/json"`
	Status *OrderResult `json:"status"`
}

// ==================== 市价平仓 ====================

type TradeMarketCloseReq struct {
	g.Meta   `path:"/trade/market-close" tags:"Trade" method:"post" summary:"市价平仓" login_required:"true"`
	WalletId int64   `json:"walletId" v:"required#请选择钱包"`
	Coin     string  `json:"coin" v:"required#请输入币种"`
	Size     float64 `json:"size"`
	Slippage float64 `json:"slippage" d:"0.05"`
}
type TradeMarketCloseRes struct {
	g.Meta `mime:"application/json"`
	Status *OrderResult `json:"status"`
}

// ==================== 限价平仓 ====================

type TradeLimitCloseReq struct {
	g.Meta   `path:"/trade/limit-close" tags:"Trade" method:"post" summary:"限价平仓" login_required:"true"`
	WalletId int64   `json:"walletId" v:"required#请选择钱包"`
	Coin     string  `json:"coin" v:"required#请输入币种"`
	Price    float64 `json:"price" v:"required#请输入价格"`
	Size     float64 `json:"size" v:"required#请输入数量"`
	Tif      string  `json:"tif" d:"Gtc"` // Gtc / Alo / Ioc
}
type TradeLimitCloseRes struct {
	g.Meta `mime:"application/json"`
	Status *OrderResult `json:"status"`
}

// ==================== 止盈/止损 ====================

type TradeSetTpSlReq struct {
	g.Meta    `path:"/trade/tpsl" tags:"Trade" method:"post" summary:"设置止盈止损" login_required:"true"`
	WalletId  int64   `json:"walletId" v:"required#请选择钱包"`
	Coin      string  `json:"coin" v:"required#请输入币种"`
	IsBuy     bool    `json:"isBuy"`
	Size      float64 `json:"size" v:"required#请输入数量"`
	TriggerPx float64 `json:"triggerPx" v:"required#请输入触发价格"`
	TpslType  string  `json:"tpslType" v:"required|in:tp,sl#请选择止盈止损类型|类型只能是tp或sl"` // tp / sl
}
type TradeSetTpSlRes struct {
	g.Meta `mime:"application/json"`
	Status *OrderResult `json:"status"`
}

// ==================== 取消订单 ====================

type TradeCancelOrderReq struct {
	g.Meta   `path:"/trade/order" tags:"Trade" method:"delete" summary:"取消订单" login_required:"true"`
	WalletId int64  `json:"walletId" v:"required#请选择钱包"`
	Coin     string `json:"coin" v:"required#请输入币种"`
	OrderId  int64  `json:"orderId" v:"required#请输入订单ID"`
}
type TradeCancelOrderRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 公共响应结构 ====================

type OrderResult struct {
	Resting *OrderResultResting `json:"resting,omitempty"`
	Filled  *OrderResultFilled  `json:"filled,omitempty"`
	Error   *string             `json:"error,omitempty"`
}

type OrderResultResting struct {
	Oid    int64  `json:"oid"`
	Status string `json:"status"`
}

type OrderResultFilled struct {
	TotalSz string `json:"totalSz"`
	AvgPx   string `json:"avgPx"`
	Oid     int    `json:"oid"`
}
