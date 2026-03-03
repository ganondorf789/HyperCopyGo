// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAppKey is the golang structure of table user_app_key for DAO operations like Where/Data.
type UserAppKey struct {
	g.Meta    `orm:"table:user_app_key, do:true"`
	Id        any         // 主键ID
	UserId    any         // 所属用户ID
	AppId     any         // AppID
	AppSecret any         // AppSecret
	Remark    any         // 备注
	ExpireAt  *gtime.Time // 过期时间,NULL表示永不过期
	Status    any         // 状态 1:启用 0:禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
