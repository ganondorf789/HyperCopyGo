// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CoinMarket is the golang structure of table coin_market for DAO operations like Where/Data.
type CoinMarket struct {
	g.Meta           `orm:"table:coin_market, do:true"`
	Id               any         // 主键ID
	Coin             any         // 币种名称
	Price            any         // 当前价格
	Change24H        any         // 24h价格变动
	ChangePercent24H any         // 24h价格变动百分比
	Open24H          any         // 24h开盘价
	Close24H         any         // 24h收盘价
	High24H          any         // 24h最高价
	Low24H           any         // 24h最低价
	Volume24H        any         // 24h成交量
	QuoteVolume24H   any         // 24h计价成交额
	Funding          any         // 资金费率
	OpenInterest     any         // 未平仓合约量
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
