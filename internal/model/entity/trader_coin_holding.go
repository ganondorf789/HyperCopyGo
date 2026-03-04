// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderCoinHolding is the golang structure for table trader_coin_holding.
type TraderCoinHolding struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`   // 主键ID
	Address   string      `json:"address"   orm:"address"    description:"钱包地址"`   // 钱包地址
	Positions string      `json:"positions" orm:"positions"  description:"持仓币种列表"` // 持仓币种列表
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`   // 更新时间
}
