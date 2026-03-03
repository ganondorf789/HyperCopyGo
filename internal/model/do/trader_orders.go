// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderOrders is the golang structure of table trader_orders for DAO operations like Where/Data.
type TraderOrders struct {
	g.Meta           `orm:"table:trader_orders, do:true"`
	Id               any         // 主键ID
	Address          any         // 钱包地址
	Coin             any         // 币种
	Side             any         // 买卖方向（A=卖/B=买）
	LimitPx          any         // 限价
	Sz               any         // 委托量
	Oid              any         // 订单ID
	Timestamp        any         // 委托时间（毫秒时间戳）
	TriggerCondition any         // 触发条件
	IsTrigger        any         // 是否触发订单
	TriggerPx        any         // 触发价
	Children         any         // 子订单(JSON)
	IsPositionTpsl   any         // 是否为仓位止盈止损
	ReduceOnly       any         // 是否只减仓
	OrderType        any         // 订单类型
	OrigSz           any         // 原始委托量
	Tif              any         // 有效期类型
	Cloid            any         // 客户端订单ID
	Status           any         // 订单状态
	CreatedAt        *gtime.Time // 创建时间
}
