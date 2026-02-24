// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyPool is the golang structure of table proxy_pool for DAO operations like Where/Data.
type ProxyPool struct {
	g.Meta    `orm:"table:proxy_pool, do:true"`
	Id        any         //
	Host      any         // 代理主机地址
	Port      any         // 代理端口
	Username  any         // 认证用户名
	Password  any         // 认证密码
	Status    any         // 状态: 1=启用 0=禁用
	Remark    any         // 备注
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
