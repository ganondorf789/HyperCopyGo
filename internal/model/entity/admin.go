// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	Username  string      `json:"username"  orm:"username"   description:""` //
	Password  string      `json:"password"  orm:"password"   description:""` //
	Realname  string      `json:"realname"  orm:"realname"   description:""` //
	Role      string      `json:"role"      orm:"role"       description:""` //
	Status    int         `json:"status"    orm:"status"     description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
