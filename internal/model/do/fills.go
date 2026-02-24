// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Fills is the golang structure of table fills for DAO operations like Where/Data.
type Fills struct {
	g.Meta        `orm:"table:fills, do:true"`
	Id            any         //
	User          any         // 交易员钱包地址
	Coin          any         // 币种
	Dir           any         // 方向描述
	Side          any         // 买卖方向 B/A
	Px            any         // 成交价格
	Sz            any         // 成交数量
	ClosedPnl     any         // 已实现盈亏
	Fee           any         // 手续费
	FeeToken      any         // 手续费币种
	BuilderFee    any         // builder手续费
	Hash          any         // 交易哈希
	Oid           any         // 订单ID
	Tid           any         // 成交ID
	Crossed       any         // 是否crossed
	StartPosition any         // 成交前持仓
	FillTime      any         // 成交时间戳(ms)
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
