// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        any         // 主键ID
	Username  any         // 用户名
	Password  any         // 密码
	Nickname  any         // 昵称
	Avatar    any         // 头像
	Email     any         // 邮箱
	Phone     any         // 手机号
	Status    any         // 状态 1:正常 0:禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
