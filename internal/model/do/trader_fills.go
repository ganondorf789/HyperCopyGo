// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderFills is the golang structure of table trader_fills for DAO operations like Where/Data.
type TraderFills struct {
	g.Meta        `orm:"table:trader_fills, do:true"`
	Id            any         // 主键ID
	Address       any         // 钱包地址
	Coin          any         // 币种
	Px            any         // 成交价
	Sz            any         // 成交量
	Side          any         // 买卖方向（A=卖/B=买）
	Time          any         // 成交时间（毫秒时间戳）
	StartPosition any         // 成交前仓位大小
	Dir           any         // 操作方向（Open Long/Open Short/Close Long/Close Short）
	ClosedPnl     any         // 平仓盈亏
	Hash          any         // 交易哈希
	Oid           any         // 订单ID
	Crossed       any         // 是否为全仓模式
	Fee           any         // 手续费
	Tid           any         // 成交ID
	Cloid         any         // 客户端订单ID
	FeeToken      any         // 手续费计价币种
	CreatedAt     *gtime.Time // 创建时间
}
