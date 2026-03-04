// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CoinMarket is the golang structure for table coin_market.
type CoinMarket struct {
	Id               int64       `json:"id"               orm:"id"                description:"主键ID"`       // 主键ID
	Coin             string      `json:"coin"             orm:"coin"              description:"币种名称"`       // 币种名称
	Price            float64     `json:"price"            orm:"price"             description:"当前价格"`       // 当前价格
	Change24H        float64     `json:"change24H"        orm:"change24h"         description:"24h价格变动"`    // 24h价格变动
	ChangePercent24H float64     `json:"changePercent24H" orm:"change_percent24h" description:"24h价格变动百分比"` // 24h价格变动百分比
	Open24H          float64     `json:"open24H"          orm:"open24h"           description:"24h开盘价"`     // 24h开盘价
	Close24H         float64     `json:"close24H"         orm:"close24h"          description:"24h收盘价"`     // 24h收盘价
	High24H          float64     `json:"high24H"          orm:"high24h"           description:"24h最高价"`     // 24h最高价
	Low24H           float64     `json:"low24H"           orm:"low24h"            description:"24h最低价"`     // 24h最低价
	Volume24H        float64     `json:"volume24H"        orm:"volume24h"         description:"24h成交量"`     // 24h成交量
	QuoteVolume24H   float64     `json:"quoteVolume24H"   orm:"quote_volume24h"   description:"24h计价成交额"`   // 24h计价成交额
	Funding          float64     `json:"funding"          orm:"funding"           description:"资金费率"`       // 资金费率
	OpenInterest     float64     `json:"openInterest"     orm:"open_interest"     description:"未平仓合约量"`     // 未平仓合约量
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`       // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:"更新时间"`       // 更新时间
}
