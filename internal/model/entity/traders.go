// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Traders is the golang structure for table traders.
type Traders struct {
	Id                     int64       `json:"id"                     orm:"id"                        description:"主键ID"`      // 主键ID
	TwitterName            string      `json:"twitterName"            orm:"twitter_name"              description:"推特显示名"`     // 推特显示名
	Username               string      `json:"username"               orm:"username"                  description:"推特用户名"`     // 推特用户名
	Address                string      `json:"address"                orm:"address"                   description:"钱包地址"`      // 钱包地址
	ProfilePicture         string      `json:"profilePicture"         orm:"profile_picture"           description:"头像链接"`      // 头像链接
	Labels                 []string    `json:"labels"                 orm:"labels"                    description:"标签列表"`      // 标签列表
	SnapEffLeverage        float64     `json:"snapEffLeverage"        orm:"snap_eff_leverage"         description:"快照-有效杠杆"`   // 快照-有效杠杆
	SnapLongPositionCount  int64       `json:"snapLongPositionCount"  orm:"snap_long_position_count"  description:"快照-多头持仓数"`  // 快照-多头持仓数
	SnapLongPositionValue  float64     `json:"snapLongPositionValue"  orm:"snap_long_position_value"  description:"快照-多头持仓价值"` // 快照-多头持仓价值
	SnapMarginUsageRate    float64     `json:"snapMarginUsageRate"    orm:"snap_margin_usage_rate"    description:"快照-保证金使用率"` // 快照-保证金使用率
	SnapPerpValue          float64     `json:"snapPerpValue"          orm:"snap_perp_value"           description:"快照-永续合约价值"` // 快照-永续合约价值
	SnapPositionCount      int64       `json:"snapPositionCount"      orm:"snap_position_count"       description:"快照-总持仓数"`   // 快照-总持仓数
	SnapPositionValue      float64     `json:"snapPositionValue"      orm:"snap_position_value"       description:"快照-总持仓价值"`  // 快照-总持仓价值
	SnapShortPositionCount int64       `json:"snapShortPositionCount" orm:"snap_short_position_count" description:"快照-空头持仓数"`  // 快照-空头持仓数
	SnapShortPositionValue float64     `json:"snapShortPositionValue" orm:"snap_short_position_value" description:"快照-空头持仓价值"` // 快照-空头持仓价值
	SnapSpotValue          float64     `json:"snapSpotValue"          orm:"snap_spot_value"           description:"快照-现货价值"`   // 快照-现货价值
	SnapTotalMarginUsed    float64     `json:"snapTotalMarginUsed"    orm:"snap_total_margin_used"    description:"快照-已用保证金"`  // 快照-已用保证金
	SnapTotalValue         float64     `json:"snapTotalValue"         orm:"snap_total_value"          description:"快照-总价值"`    // 快照-总价值
	SnapUnrealizedPnl      float64     `json:"snapUnrealizedPnl"      orm:"snap_unrealized_pnl"       description:"快照-未实现盈亏"`  // 快照-未实现盈亏
	ShortPnl               float64     `json:"shortPnl"               orm:"short_pnl"                 description:"空头盈亏"`      // 空头盈亏
	ShortWinRate           float64     `json:"shortWinRate"           orm:"short_win_rate"            description:"空头胜率"`      // 空头胜率
	LongPnl                float64     `json:"longPnl"                orm:"long_pnl"                  description:"多头盈亏"`      // 多头盈亏
	LongWinRate            float64     `json:"longWinRate"            orm:"long_win_rate"             description:"多头胜率"`      // 多头胜率
	CreatedAt              *gtime.Time `json:"createdAt"              orm:"created_at"                description:"创建时间"`      // 创建时间
	UpdatedAt              *gtime.Time `json:"updatedAt"              orm:"updated_at"                description:"更新时间"`      // 更新时间
}
