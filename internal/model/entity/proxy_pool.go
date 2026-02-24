// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyPool is the golang structure for table proxy_pool.
type ProxyPool struct {
	Id        int64       `json:"id"        orm:"id"         description:""`              //
	Host      string      `json:"host"      orm:"host"       description:"代理主机地址"`        // 代理主机地址
	Port      int         `json:"port"      orm:"port"       description:"代理端口"`          // 代理端口
	Username  string      `json:"username"  orm:"username"   description:"认证用户名"`         // 认证用户名
	Password  string      `json:"password"  orm:"password"   description:"认证密码"`          // 认证密码
	Status    int         `json:"status"    orm:"status"     description:"状态: 1=启用 0=禁用"` // 状态: 1=启用 0=禁用
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`            // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`              //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`              //
}
