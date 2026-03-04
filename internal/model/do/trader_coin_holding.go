// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderCoinHolding is the golang structure of table trader_coin_holding for DAO operations like Where/Data.
type TraderCoinHolding struct {
	g.Meta    `orm:"table:trader_coin_holding, do:true"`
	Id        any         // 主键ID
	Address   any         // 钱包地址
	Positions any         // 持仓币种列表
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
