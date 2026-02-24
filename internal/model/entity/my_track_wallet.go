// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MyTrackWallet is the golang structure for table my_track_wallet.
type MyTrackWallet struct {
	Id           int64       `json:"id"           orm:"id"            description:""`                         //
	UserId       int64       `json:"userId"       orm:"user_id"       description:"所属用户ID"`                   // 所属用户ID
	Wallet       string      `json:"wallet"       orm:"wallet"        description:"跟踪的钱包地址"`                  // 跟踪的钱包地址
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`                       // 备注
	EnableNotify int         `json:"enableNotify" orm:"enable_notify" description:"是否开启通知 0:关 1:开"`           // 是否开启通知 0:关 1:开
	NotifyAction string      `json:"notifyAction" orm:"notify_action" description:"通知动作 1:开仓 2:平仓 3:加仓 4:减仓"` // 通知动作 1:开仓 2:平仓 3:加仓 4:减仓
	Lang         string      `json:"lang"         orm:"lang"          description:"语言"`                       // 语言
	Status       int         `json:"status"       orm:"status"        description:"状态 1:正常 0:禁用"`             // 状态 1:正常 0:禁用
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`                         //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`                         //
}
