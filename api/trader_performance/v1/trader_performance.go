package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 交易绩效统计
type TraderPerformanceReq struct {
	g.Meta  `path:"/trader-performance" tags:"TraderPerformance" method:"get" summary:"交易绩效统计" login_required:"true"`
	Address string `json:"address" v:"required#钱包地址不能为空"`
	Window  string `json:"window" d:"allTime" v:"in:day,week,month,allTime#时间窗口不合法"`
}
type TraderPerformanceRes struct {
	g.Meta            `mime:"application/json"`
	TotalPnl          float64                           `json:"totalPnl"`
	Gross             float64                           `json:"gross"`
	Fees              float64                           `json:"fees"`
	LongPnl           float64                           `json:"longPnl"`
	ShortPnl          float64                           `json:"shortPnl"`
	WinRate           float64                           `json:"winRate"`
	Winning           int                               `json:"winning"`
	Total             int                               `json:"total"`
	TradeDuration     int64                             `json:"tradeDuration"`
	MinDuration       int64                             `json:"minDuration"`
	MaxDuration       int64                             `json:"maxDuration"`
	BestTrades        []model.TradePerformanceBestTrade `json:"bestTrades"`
	PerformanceAssets []model.TradePerformanceAsset     `json:"performanceAssets"`
}

// 交易概览
type TraderPerformanceSummaryReq struct {
	g.Meta  `path:"/trader-performance/summary" tags:"TraderPerformance" method:"get" summary:"交易概览" login_required:"true"`
	Address string `json:"address" v:"required#钱包地址不能为空"`
}
type TraderPerformanceSummaryRes struct {
	g.Meta         `mime:"application/json"`
	WinRate        float64 `json:"winRate"`
	OrderCount     int     `json:"orderCount"`
	ClosePosCount  int     `json:"closePosCount"`
	AvgPosDuration int64   `json:"avgPosDuration"`
	MaxDrawdown    float64 `json:"maxDrawdown"`
	TotalPnl       float64 `json:"totalPnl"`
}
