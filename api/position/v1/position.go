package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取持仓列表（无需登录）
type PositionListReq struct {
	g.Meta   `path:"/position" tags:"Position" method:"get" summary:"持仓列表"`
	User     string `json:"user" in:"query"`   // 按钱包地址筛选
	Symbol   string `json:"symbol" in:"query"` // 按交易对筛选
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type PositionListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.PositionItem `json:"list"`
	Total  int                  `json:"total"`
	Page   int                  `json:"page"`
}

// 获取持仓统计（无需登录）
type PositionStatsReq struct {
	g.Meta   `path:"/position/stats" tags:"Position" method:"get" summary:"持仓统计"`
	Interval string `json:"interval" in:"query" v:"required|in:5m,30m,1h,4h,12h,1D#请输入时间框架|时间框架仅支持5m,30m,1h,4h,12h,1D"` // 时间框架
	Symbol   string `json:"symbol" in:"query"`                                                                                // 按交易对筛选
}
type PositionStatsRes struct {
	g.Meta `mime:"application/json"`
	List   []model.PositionStatsPoint `json:"list"`
}

// 多空比率曲线（无需登录）
type PositionLongShortRatioReq struct {
	g.Meta   `path:"/position/long-short-ratio" tags:"Position" method:"get" summary:"多空比率曲线"`
	Interval string `json:"interval" in:"query" v:"required|in:1h,4h,1D#请输入时间框架|时间框架仅支持1h,4h,1D"` // 时间框架
	Symbol   string `json:"symbol" in:"query"`                                                        // 按交易对筛选
}
type PositionLongShortRatioRes struct {
	g.Meta `mime:"application/json"`
	List   []model.LongShortRatioPoint `json:"list"`
}
