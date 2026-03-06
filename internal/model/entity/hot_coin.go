// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HotCoin is the golang structure for table hot_coin.
type HotCoin struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"` // 主键ID
	Coin      string      `json:"coin"      orm:"coin"       description:"币种名称"` // 币种名称
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
}
