// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderOrders is the golang structure for table trader_orders.
type TraderOrders struct {
	Id               int64       `json:"id"               orm:"id"                description:"主键ID"`          // 主键ID
	Address          string      `json:"address"          orm:"address"           description:"钱包地址"`          // 钱包地址
	Coin             string      `json:"coin"             orm:"coin"              description:"币种"`            // 币种
	Side             string      `json:"side"             orm:"side"              description:"买卖方向（A=卖/B=买）"` // 买卖方向（A=卖/B=买）
	LimitPx          float64     `json:"limitPx"          orm:"limit_px"          description:"限价"`            // 限价
	Sz               float64     `json:"sz"               orm:"sz"                description:"委托量"`           // 委托量
	Oid              int64       `json:"oid"              orm:"oid"               description:"订单ID"`          // 订单ID
	Timestamp        int64       `json:"timestamp"        orm:"timestamp"         description:"委托时间（毫秒时间戳）"`   // 委托时间（毫秒时间戳）
	TriggerCondition string      `json:"triggerCondition" orm:"trigger_condition" description:"触发条件"`          // 触发条件
	IsTrigger        bool        `json:"isTrigger"        orm:"is_trigger"        description:"是否触发订单"`        // 是否触发订单
	TriggerPx        float64     `json:"triggerPx"        orm:"trigger_px"        description:"触发价"`           // 触发价
	Children         string      `json:"children"         orm:"children"          description:"子订单(JSON)"`     // 子订单(JSON)
	IsPositionTpsl   bool        `json:"isPositionTpsl"   orm:"is_position_tpsl"  description:"是否为仓位止盈止损"`     // 是否为仓位止盈止损
	ReduceOnly       bool        `json:"reduceOnly"       orm:"reduce_only"       description:"是否只减仓"`         // 是否只减仓
	OrderType        string      `json:"orderType"        orm:"order_type"        description:"订单类型"`          // 订单类型
	OrigSz           float64     `json:"origSz"           orm:"orig_sz"           description:"原始委托量"`         // 原始委托量
	Tif              string      `json:"tif"              orm:"tif"               description:"有效期类型"`         // 有效期类型
	Cloid            string      `json:"cloid"            orm:"cloid"             description:"客户端订单ID"`       // 客户端订单ID
	Status           string      `json:"status"           orm:"status"            description:"订单状态"`          // 订单状态
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`          // 创建时间
}
