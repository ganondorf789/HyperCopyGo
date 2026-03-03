// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WhaleAnchor is the golang structure for table whale_anchor.
type WhaleAnchor struct {
	Id             int64       `json:"id"             orm:"id"              description:"主键ID"`                                    // 主键ID
	Symbol         string      `json:"symbol"         orm:"symbol"          description:"交易对符号"`                                   // 交易对符号
	Volume24H      float64     `json:"volume24H"      orm:"volume24h"       description:"24h成交量(USD)"`                             // 24h成交量(USD)
	OpenInterest   float64     `json:"openInterest"   orm:"open_interest"   description:"当前未平仓合约量(USD)"`                           // 当前未平仓合约量(USD)
	Depth1Pct      float64     `json:"depth1Pct"      orm:"depth1pct"       description:"1%盘口深度(USD)"`                             // 1%盘口深度(USD)
	ValVolume      float64     `json:"valVolume"      orm:"val_volume"      description:"0.4% x 24h Volume"`                       // 0.4% x 24h Volume
	ValOi          float64     `json:"valOi"          orm:"val_oi"          description:"1% x OI"`                                 // 1% x OI
	ValDepth       float64     `json:"valDepth"       orm:"val_depth"       description:"30% x 1% Depth"`                          // 30% x 1% Depth
	WhaleThreshold float64     `json:"whaleThreshold" orm:"whale_threshold" description:"巨鲸仓位阈值 max(val_volume,val_oi,val_depth)"` // 巨鲸仓位阈值 max(val_volume,val_oi,val_depth)
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`                                    // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`                                    // 更新时间
}
