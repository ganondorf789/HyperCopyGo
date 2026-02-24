// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Fills is the golang structure for table fills.
type Fills struct {
	Id            int64       `json:"id"            orm:"id"             description:""`           //
	User          string      `json:"user"          orm:"user"           description:"交易员钱包地址"`    // 交易员钱包地址
	Coin          string      `json:"coin"          orm:"coin"           description:"币种"`         // 币种
	Dir           string      `json:"dir"           orm:"dir"            description:"方向描述"`       // 方向描述
	Side          string      `json:"side"          orm:"side"           description:"买卖方向 B/A"`   // 买卖方向 B/A
	Px            float64     `json:"px"            orm:"px"             description:"成交价格"`       // 成交价格
	Sz            float64     `json:"sz"            orm:"sz"             description:"成交数量"`       // 成交数量
	ClosedPnl     float64     `json:"closedPnl"     orm:"closed_pnl"     description:"已实现盈亏"`      // 已实现盈亏
	Fee           float64     `json:"fee"           orm:"fee"            description:"手续费"`        // 手续费
	FeeToken      string      `json:"feeToken"      orm:"fee_token"      description:"手续费币种"`      // 手续费币种
	BuilderFee    float64     `json:"builderFee"    orm:"builder_fee"    description:"builder手续费"` // builder手续费
	Hash          string      `json:"hash"          orm:"hash"           description:"交易哈希"`       // 交易哈希
	Oid           int64       `json:"oid"           orm:"oid"            description:"订单ID"`       // 订单ID
	Tid           int64       `json:"tid"           orm:"tid"            description:"成交ID"`       // 成交ID
	Crossed       bool        `json:"crossed"       orm:"crossed"        description:"是否crossed"`  // 是否crossed
	StartPosition float64     `json:"startPosition" orm:"start_position" description:"成交前持仓"`      // 成交前持仓
	FillTime      int64       `json:"fillTime"      orm:"fill_time"      description:"成交时间戳(ms)"`  // 成交时间戳(ms)
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`           //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`           //
}
