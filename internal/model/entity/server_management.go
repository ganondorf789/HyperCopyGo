// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ServerManagement is the golang structure for table server_management.
type ServerManagement struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`    // 主键ID
	Ip        string      `json:"ip"        orm:"ip"         description:"服务器IP地址"` // 服务器IP地址
	Username  string      `json:"username"  orm:"username"   description:"用户名"`     // 用户名
	Password  string      `json:"password"  orm:"password"   description:"密码"`      // 密码
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`    // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`    // 更新时间
}
