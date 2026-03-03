package model

import "github.com/gogf/gf/v2/os/gtime"

// CompletedTradeItem 已完成交易列表项
type CompletedTradeItem struct {
	Id         int64       `json:"id"`
	Address    string      `json:"address"`
	Coin       string      `json:"coin"`
	MarginMode string      `json:"marginMode"`
	Direction  string      `json:"direction"`
	Size       float64     `json:"size"`
	EntryPrice float64     `json:"entryPrice"`
	ClosePrice float64     `json:"closePrice"`
	StartTime  int64       `json:"startTime"`
	EndTime    int64       `json:"endTime"`
	TotalFee   float64     `json:"totalFee"`
	Pnl        float64     `json:"pnl"`
	FillCount  int64       `json:"fillCount"`
	CreatedAt  *gtime.Time `json:"createdAt"`
	UpdatedAt  *gtime.Time `json:"updatedAt"`
}

type TradePerformanceBestTrade struct {
	Coin      string  `json:"coin"`
	Direction string  `json:"direction"`
	Duration  int64   `json:"duration"`
	CreateAt  int64   `json:"createAt"`
	Pnl       float64 `json:"pnl"`
}

type TradePerformanceAsset struct {
	Address *string `json:"address"`
	Trades  int     `json:"trades"`
	Coin    string  `json:"coin"`
	Pnl     float64 `json:"pnl"`
	Fees    float64 `json:"fees"`
}
