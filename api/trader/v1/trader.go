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
