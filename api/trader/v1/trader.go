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
