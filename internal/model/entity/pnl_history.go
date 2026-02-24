// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PnlHistory is the golang structure for table pnl_history.
type PnlHistory struct {
	Id        int64       `json:"id"        orm:"id"         description:""`                       //
	User      string      `json:"user"      orm:"user"       description:"交易员钱包地址"`                // 交易员钱包地址
	Timeframe string      `json:"timeframe" orm:"timeframe"  description:"时间框架: 1D, 7D, 30D, All"` // 时间框架: 1D, 7D, 30D, All
	PnlList   string      `json:"pnlList"   orm:"pnl_list"   description:"PnL数据点列表"`               // PnL数据点列表
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                       //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`                       //
}
