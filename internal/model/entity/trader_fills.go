// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TraderFills is the golang structure for table trader_fills.
type TraderFills struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键ID"`                                              // 主键ID
	Address       string      `json:"address"       orm:"address"        description:"钱包地址"`                                              // 钱包地址
	Coin          string      `json:"coin"          orm:"coin"           description:"币种"`                                                // 币种
	Px            float64     `json:"px"            orm:"px"             description:"成交价"`                                               // 成交价
	Sz            float64     `json:"sz"            orm:"sz"             description:"成交量"`                                               // 成交量
	Side          string      `json:"side"          orm:"side"           description:"买卖方向（A=卖/B=买）"`                                     // 买卖方向（A=卖/B=买）
	Time          int64       `json:"time"          orm:"time"           description:"成交时间（毫秒时间戳）"`                                       // 成交时间（毫秒时间戳）
	StartPosition float64     `json:"startPosition" orm:"start_position" description:"成交前仓位大小"`                                           // 成交前仓位大小
	Dir           string      `json:"dir"           orm:"dir"            description:"操作方向（Open Long/Open Short/Close Long/Close Short）"` // 操作方向（Open Long/Open Short/Close Long/Close Short）
	ClosedPnl     float64     `json:"closedPnl"     orm:"closed_pnl"     description:"平仓盈亏"`                                              // 平仓盈亏
	Hash          string      `json:"hash"          orm:"hash"           description:"交易哈希"`                                              // 交易哈希
	Oid           int64       `json:"oid"           orm:"oid"            description:"订单ID"`                                              // 订单ID
	Crossed       bool        `json:"crossed"       orm:"crossed"        description:"是否为全仓模式"`                                           // 是否为全仓模式
	Fee           float64     `json:"fee"           orm:"fee"            description:"手续费"`                                               // 手续费
	Tid           int64       `json:"tid"           orm:"tid"            description:"成交ID"`                                              // 成交ID
	Cloid         string      `json:"cloid"         orm:"cloid"          description:"客户端订单ID"`                                           // 客户端订单ID
	FeeToken      string      `json:"feeToken"      orm:"fee_token"      description:"手续费计价币种"`                                           // 手续费计价币种
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`                                              // 创建时间
}
