// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MyTrackWallet is the golang structure of table my_track_wallet for DAO operations like Where/Data.
type MyTrackWallet struct {
	g.Meta       `orm:"table:my_track_wallet, do:true"`
	Id           any         //
	UserId       any         // 所属用户ID
	Wallet       any         // 跟踪的钱包地址
	Remark       any         // 备注
	EnableNotify any         // 是否开启通知 0:关 1:开
	NotifyAction any         // 通知动作 1:开仓 2:平仓 3:加仓 4:减仓
	Lang         any         // 语言
	Status       any         // 状态 1:正常 0:禁用
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
