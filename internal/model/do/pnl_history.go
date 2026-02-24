// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PnlHistory is the golang structure of table pnl_history for DAO operations like Where/Data.
type PnlHistory struct {
	g.Meta    `orm:"table:pnl_history, do:true"`
	Id        any         //
	User      any         // 交易员钱包地址
	Timeframe any         // 时间框架: 1D, 7D, 30D, All
	PnlList   any         // PnL数据点列表
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
