// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NotificationRead is the golang structure for table notification_read.
type NotificationRead struct {
	Id             int64       `json:"id"             orm:"id"              description:"主键ID"` // 主键ID
	UserId         int64       `json:"userId"         orm:"user_id"         description:"用户ID"` // 用户ID
	NotificationId int64       `json:"notificationId" orm:"notification_id" description:"通知ID"` // 通知ID
	ReadAt         *gtime.Time `json:"readAt"         orm:"read_at"         description:"已读时间"` // 已读时间
}
