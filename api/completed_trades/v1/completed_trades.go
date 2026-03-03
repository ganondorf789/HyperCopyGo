package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 已完成交易列表
type CompletedTradesListReq struct {
	g.Meta    `path:"/completed-trades" tags:"CompletedTrades" method:"get" summary:"已完成交易列表" login_required:"true"`
	Address   string `json:"address"`
	Coin      string `json:"coin"`
	Direction string `json:"direction"`
	Window    string `json:"window" d:"allTime" v:"in:day,week,month,allTime#时间窗口不合法"`
	Page      int    `json:"page" d:"1"`
	PageSize  int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type CompletedTradesListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.CompletedTradeItem `json:"list"`
	Total  int                        `json:"total"`
	Page   int                        `json:"page"`
}

