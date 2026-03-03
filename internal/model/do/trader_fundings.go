// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderFundings is the golang structure of table trader_fundings for DAO operations like Where/Data.
type TraderFundings struct {
	g.Meta      `orm:"table:trader_fundings, do:true"`
	Id          any         // 主键ID
	Address     any         // 钱包地址
	Time        any         // 资金费时间（毫秒时间戳）
	Hash        any         // 交易哈希
	Coin        any         // 币种
	Usdc        any         // USDC金额（正=收入，负=支出）
	Szi         any         // 持仓大小
	FundingRate any         // 资金费率
	CreatedAt   *gtime.Time // 创建时间
}
