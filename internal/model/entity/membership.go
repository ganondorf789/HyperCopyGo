// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Membership is the golang structure for table membership.
type Membership struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`                     // 主键ID
	UserId    int64       `json:"userId"    orm:"user_id"    description:"所属用户ID"`                   // 所属用户ID
	Level     int         `json:"level"     orm:"level"      description:"会员等级 0:免费 1:基础 2:高级 3:专业"` // 会员等级 0:免费 1:基础 2:高级 3:专业
	StartAt   *gtime.Time `json:"startAt"   orm:"start_at"   description:"会员开始时间"`                   // 会员开始时间
	ExpireAt  *gtime.Time `json:"expireAt"  orm:"expire_at"  description:"会员到期时间"`                   // 会员到期时间
	Status    int         `json:"status"    orm:"status"     description:"状态 1:正常 0:禁用"`             // 状态 1:正常 0:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                     // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                     // 更新时间
}
