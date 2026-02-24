// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	Username  string      `json:"username"  orm:"username"   description:""` //
	Password  string      `json:"password"  orm:"password"   description:""` //
	Nickname  string      `json:"nickname"  orm:"nickname"   description:""` //
	Avatar    string      `json:"avatar"    orm:"avatar"     description:""` //
	Email     string      `json:"email"     orm:"email"      description:""` //
	Phone     string      `json:"phone"     orm:"phone"      description:""` //
	Status    int         `json:"status"    orm:"status"     description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
