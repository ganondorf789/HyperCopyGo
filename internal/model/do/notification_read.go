// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NotificationRead is the golang structure of table notification_read for DAO operations like Where/Data.
type NotificationRead struct {
	g.Meta         `orm:"table:notification_read, do:true"`
	Id             any         // 主键ID
	UserId         any         // 用户ID
	NotificationId any         // 通知ID
	ReadAt         *gtime.Time // 已读时间
}
