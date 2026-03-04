// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Traders is the golang structure of table traders for DAO operations like Where/Data.
type Traders struct {
	g.Meta                 `orm:"table:traders, do:true"`
	Id                     any         // 主键ID
	TwitterName            any         // 推特显示名
	Username               any         // 推特用户名
	Address                any         // 钱包地址
	ProfilePicture         any         // 头像链接
	IsHotAddress           any         // 是否热门地址
	IsTwitterKol           any         // 是否推特KOL
	Labels                 []string    // 标签列表
	SnapEffLeverage        any         // 快照-有效杠杆
	SnapLongPositionCount  any         // 快照-多头持仓数
	SnapLongPositionValue  any         // 快照-多头持仓价值
	SnapMarginUsageRate    any         // 快照-保证金使用率
	SnapPerpValue          any         // 快照-永续合约价值
	SnapPositionCount      any         // 快照-总持仓数
	SnapPositionValue      any         // 快照-总持仓价值
	SnapShortPositionCount any         // 快照-空头持仓数
	SnapShortPositionValue any         // 快照-空头持仓价值
	SnapSpotValue          any         // 快照-现货价值
	SnapTotalMarginUsed    any         // 快照-已用保证金
	SnapTotalValue         any         // 快照-总价值
	SnapUnrealizedPnl      any         // 快照-未实现盈亏
	ShortPnl               any         // 空头盈亏
	ShortWinRate           any         // 空头胜率
	LongPnl                any         // 多头盈亏
	LongWinRate            any         // 多头胜率
	TotalPnl               any         // 总盈亏
	CreatedAt              *gtime.Time // 创建时间
	UpdatedAt              *gtime.Time // 更新时间
}
