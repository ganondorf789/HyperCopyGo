package model

import "github.com/gogf/gf/v2/os/gtime"

// TraderPositionItem 交易员持仓列表项
type TraderPositionItem struct {
	Id                    int64       `json:"id"`
	Address               string      `json:"address"`
	Coin                  string      `json:"coin"`
	Szi                   float64     `json:"szi"`
	LeverageType          string      `json:"leverageType"`
	Leverage              int64       `json:"leverage"`
	EntryPx               float64     `json:"entryPx"`
	PositionValue         float64     `json:"positionValue"`
	UnrealizedPnl         float64     `json:"unrealizedPnl"`
	ReturnOnEquity        float64     `json:"returnOnEquity"`
	LiquidationPx         float64     `json:"liquidationPx"`
	MarginUsed            float64     `json:"marginUsed"`
	MaxLeverage           int64       `json:"maxLeverage"`
	CumFundingAllTime     float64     `json:"cumFundingAllTime"`
	CumFundingSinceOpen   float64     `json:"cumFundingSinceOpen"`
	CumFundingSinceChange float64     `json:"cumFundingSinceChange"`
	CreatedAt             *gtime.Time `json:"createdAt"`
	UpdatedAt             *gtime.Time `json:"updatedAt"`
}
