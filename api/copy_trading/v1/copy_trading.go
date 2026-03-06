package v1

import (
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CopyTradingListReq struct {
	g.Meta   `path:"/copy-trading-orders" tags:"CopyTrading" method:"get" summary:"跟单交易列表" login_required:"true"`
	Status   string `json:"status" in:"query"`
	Page     int    `json:"page" in:"query" d:"1"`
	PageSize int    `json:"pageSize" in:"query" d:"20" v:"max:100#每页最多100条"`
}
type CopyTradingListRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.CopyTrading `json:"list"`
	Total  int                  `json:"total"`
	Page   int                  `json:"page"`
}

type CopyTradingStopReq struct {
	g.Meta `path:"/copy-trading-orders/{id}/stop" tags:"CopyTrading" method:"post" summary:"停止跟单" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type CopyTradingStopRes struct {
	g.Meta `mime:"application/json"`
}
