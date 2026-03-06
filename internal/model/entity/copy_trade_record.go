// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CopyTradeRecord is the golang structure for table copy_trade_record.
type CopyTradeRecord struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键ID"`                                            // 主键ID
	UserId        int64       `json:"userId"        orm:"user_id"        description:"所属用户ID"`                                          // 所属用户ID
	Address       string      `json:"address"       orm:"address"        description:"钱包地址"`                                            // 钱包地址
	Coin          string      `json:"coin"          orm:"coin"           description:"币种"`                                              // 币种
	Direction     string      `json:"direction"     orm:"direction"      description:"方向（Open Long/Open Short/Close Long/Close Short）"` // 方向（Open Long/Open Short/Close Long/Close Short）
	Size          float64     `json:"size"          orm:"size"           description:"成交规模（张数）"`                                        // 成交规模（张数）
	Price         float64     `json:"price"         orm:"price"          description:"成交价格"`                                            // 成交价格
	ClosedPnl     float64     `json:"closedPnl"     orm:"closed_pnl"     description:"已实现盈亏（USD）"`                                      // 已实现盈亏（USD）
	ExecuteStatus int64       `json:"executeStatus" orm:"execute_status" description:"执行状态 0:待执行 1:成功 2:失败 3:跳过"`                       // 执行状态 0:待执行 1:成功 2:失败 3:跳过
	OrderStatus   string      `json:"orderStatus"   orm:"order_status"   description:"订单状态 open/filled/canceled/triggered"`             // 订单状态 open/filled/canceled/triggered
	ErrorMsg      string      `json:"errorMsg"      orm:"error_msg"      description:"执行失败原因"`                                          // 执行失败原因
	TradeTime     *gtime.Time `json:"tradeTime"     orm:"trade_time"     description:"触发交易时间（源头成交时间）"`                                  // 触发交易时间（源头成交时间）
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`                                            // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`                                            // 更新时间
}
