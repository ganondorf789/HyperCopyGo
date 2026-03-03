// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Membership is the golang structure of table membership for DAO operations like Where/Data.
type Membership struct {
	g.Meta    `orm:"table:membership, do:true"`
	Id        any         // 主键ID
	UserId    any         // 所属用户ID
	Level     any         // 会员等级 0:免费 1:基础 2:高级 3:专业
	StartAt   *gtime.Time // 会员开始时间
	ExpireAt  *gtime.Time // 会员到期时间
	Status    any         // 状态 1:正常 0:禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
