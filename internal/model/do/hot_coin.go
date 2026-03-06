// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HotCoin is the golang structure of table hot_coin for DAO operations like Where/Data.
type HotCoin struct {
	g.Meta    `orm:"table:hot_coin, do:true"`
	Id        any         // 主键ID
	Coin      any         // 币种名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
