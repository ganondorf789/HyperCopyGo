// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhaleAnchor is the golang structure of table whale_anchor for DAO operations like Where/Data.
type WhaleAnchor struct {
	g.Meta         `orm:"table:whale_anchor, do:true"`
	Id             any         // 主键ID
	Symbol         any         // 交易对符号
	Volume24H      any         // 24h成交量(USD)
	OpenInterest   any         // 当前未平仓合约量(USD)
	Depth1Pct      any         // 1%盘口深度(USD)
	ValVolume      any         // 0.4% x 24h Volume
	ValOi          any         // 1% x OI
	ValDepth       any         // 30% x 1% Depth
	WhaleThreshold any         // 巨鲸仓位阈值 max(val_volume,val_oi,val_depth)
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
