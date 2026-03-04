// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Leaderboard is the golang structure for table leaderboard.
type Leaderboard struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键ID"`  // 主键ID
	EthAddress   string      `json:"ethAddress"   orm:"eth_address"   description:"钱包地址"`  // 钱包地址
	AccountValue float64     `json:"accountValue" orm:"account_value" description:"账户价值"`  // 账户价值
	Pnl          float64     `json:"pnl"          orm:"pnl"           description:"盈亏"`    // 盈亏
	Roi          float64     `json:"roi"          orm:"roi"           description:"投资回报率"` // 投资回报率
	Vlm          float64     `json:"vlm"          orm:"vlm"           description:"交易量"`   // 交易量
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`  // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`  // 更新时间
}
