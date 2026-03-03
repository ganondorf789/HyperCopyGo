// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderFundings is the golang structure for table trader_fundings.
type TraderFundings struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键ID"`              // 主键ID
	Address     string      `json:"address"     orm:"address"      description:"钱包地址"`              // 钱包地址
	Time        int64       `json:"time"        orm:"time"         description:"资金费时间（毫秒时间戳）"`      // 资金费时间（毫秒时间戳）
	Hash        string      `json:"hash"        orm:"hash"         description:"交易哈希"`              // 交易哈希
	Coin        string      `json:"coin"        orm:"coin"         description:"币种"`                // 币种
	Usdc        float64     `json:"usdc"        orm:"usdc"         description:"USDC金额（正=收入，负=支出）"` // USDC金额（正=收入，负=支出）
	Szi         float64     `json:"szi"         orm:"szi"          description:"持仓大小"`              // 持仓大小
	FundingRate float64     `json:"fundingRate" orm:"funding_rate" description:"资金费率"`              // 资金费率
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`              // 创建时间
}
