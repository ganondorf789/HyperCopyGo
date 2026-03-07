// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ServerManagement is the golang structure of table server_management for DAO operations like Where/Data.
type ServerManagement struct {
	g.Meta    `orm:"table:server_management, do:true"`
	Id        any         // 主键ID
	Ip        any         // 服务器IP地址
	Username  any         // 用户名
	Password  any         // 密码
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
