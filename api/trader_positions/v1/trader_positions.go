package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 交易员持仓列表
type TraderPositionsListReq struct {
	g.Meta    `path:"/trader-positions" tags:"TraderPositions" method:"get" summary:"交易员持仓列表" login_required:"true"`
	Address   string `json:"address"`
	Coin      string `json:"coin"`
	Direction string `json:"direction"`
	Page      int    `json:"page" d:"1"`
	PageSize  int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type TraderPositionsListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.TraderPositionItem `json:"list"`
	Total  int                        `json:"total"`
	Page   int                        `json:"page"`
}
