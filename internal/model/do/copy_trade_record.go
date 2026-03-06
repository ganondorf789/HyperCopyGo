// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CopyTradeRecord is the golang structure of table copy_trade_record for DAO operations like Where/Data.
type CopyTradeRecord struct {
	g.Meta        `orm:"table:copy_trade_record, do:true"`
	Id            any         // 主键ID
	UserId        any         // 所属用户ID
	Address       any         // 钱包地址
	Coin          any         // 币种
	Direction     any         // 方向（Open Long/Open Short/Close Long/Close Short）
	Size          any         // 成交规模（张数）
	Price         any         // 成交价格
	ClosedPnl     any         // 已实现盈亏（USD）
	ExecuteStatus any         // 执行状态 0:待执行 1:成功 2:失败 3:跳过
	OrderStatus   any         // 订单状态 open/filled/canceled/triggered
	ErrorMsg      any         // 执行失败原因
	TradeTime     *gtime.Time // 触发交易时间（源头成交时间）
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
