// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notification is the golang structure of table notification for DAO operations like Where/Data.
type Notification struct {
	g.Meta    `orm:"table:notification, do:true"`
	Id        any         // 主键ID
	UserId    any         // 所属用户ID,0表示公共通知
	Category  any         // 通知类型 public/copy_trading/whale/track/market
	Title     any         // 通知标题
	Content   any         // 通知内容
	RefId     any         // 关联业务ID,0表示无关联
	RefType   any         // 关联业务类型 copy_trading/track_wallet/position
	Level     any         // 通知级别 0:普通 1:重要 2:紧急
	Status    any         // 状态 1:正常 0:已撤回
	CreatedAt *gtime.Time // 创建时间
}
