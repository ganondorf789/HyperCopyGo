// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAppKey is the golang structure for table user_app_key.
type UserAppKey struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`            // 主键ID
	UserId    int64       `json:"userId"    orm:"user_id"    description:"所属用户ID"`          // 所属用户ID
	AppId     string      `json:"appId"     orm:"app_id"     description:"AppID"`           // AppID
	AppSecret string      `json:"appSecret" orm:"app_secret" description:"AppSecret"`       // AppSecret
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`              // 备注
	ExpireAt  *gtime.Time `json:"expireAt"  orm:"expire_at"  description:"过期时间,NULL表示永不过期"` // 过期时间,NULL表示永不过期
	Status    int         `json:"status"    orm:"status"     description:"状态 1:启用 0:禁用"`    // 状态 1:启用 0:禁用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`            // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`            // 更新时间
}
