package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 热门地址列表（免鉴权）
type TraderPopularReq struct {
	g.Meta `path:"/trader/popular" tags:"Trader" method:"get" summary:"获取热门地址列表"`
}
type TraderPopularRes struct {
	g.Meta `mime:"application/json"`
	List   []model.PopularTraderItem `json:"list"`
}

// Discover 发现交易员（免鉴权，支持排序/筛选/标签/币种/分页）
type TraderDiscoverReq struct {
	g.Meta   `path:"/trader/discover" tags:"Trader" method:"post" summary:"发现交易员"`
	Window   string                `json:"window" v:"required|in:day,week,month,allTime#窗口为必填项|窗口仅支持day/week/month/allTime"`
	Sort     *model.DiscoverSort   `json:"sort"`
	Filters  []model.DiscoverFilter `json:"filters"`
	Coins    []string              `json:"coins"`
	Tags     []string              `json:"tags"`
	Page     int                   `json:"page" d:"1"`
	PageSize int                   `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type TraderDiscoverRes struct {
	g.Meta `mime:"application/json"`
	List   []model.DiscoverTraderItem `json:"list"`
	Total  int                        `json:"total"`
	Page   int                        `json:"page"`
}

// X KOL 列表（免鉴权，分页 + 窗口筛选）
type TraderKolListReq struct {
	g.Meta   `path:"/trader/kol" tags:"Trader" method:"get" summary:"获取X KOL列表"`
	Window   string `json:"window" d:"month" v:"in:day,week,month,allTime#窗口参数仅支持day/week/month/allTime"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type TraderKolListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.KolTraderItem `json:"list"`
	Total  int                   `json:"total"`
	Page   int                   `json:"page"`
}
