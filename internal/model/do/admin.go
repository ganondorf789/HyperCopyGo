// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:admin, do:true"`
	Id        any         // 主键ID
	Username  any         // 用户名
	Password  any         // 密码
	Realname  any         // 真实姓名
	Role      any         // 角色 admin/super_admin
	Status    any         // 状态 1:正常 0:禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
