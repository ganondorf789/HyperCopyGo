// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        int64       `json:"id"        orm:"id"         description:""`                     //
	Username  string      `json:"username"  orm:"username"   description:"用户名"`                  // 用户名
	Password  string      `json:"password"  orm:"password"   description:"密码"`                   // 密码
	Realname  string      `json:"realname"  orm:"realname"   description:"真实姓名"`                 // 真实姓名
	Role      string      `json:"role"      orm:"role"       description:"角色 admin/super_admin"` // 角色 admin/super_admin
	Status    int         `json:"status"    orm:"status"     description:"状态 1:正常 0:禁用"`         // 状态 1:正常 0:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                     //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`                     //
}
