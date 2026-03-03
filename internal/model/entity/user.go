// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`         // 主键ID
	Username  string      `json:"username"  orm:"username"   description:"用户名"`          // 用户名
	Password  string      `json:"password"  orm:"password"   description:"密码"`           // 密码
	Nickname  string      `json:"nickname"  orm:"nickname"   description:"昵称"`           // 昵称
	Avatar    string      `json:"avatar"    orm:"avatar"     description:"头像"`           // 头像
	Email     string      `json:"email"     orm:"email"      description:"邮箱"`           // 邮箱
	Phone     string      `json:"phone"     orm:"phone"      description:"手机号"`          // 手机号
	Status    int         `json:"status"    orm:"status"     description:"状态 1:正常 0:禁用"` // 状态 1:正常 0:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
}
