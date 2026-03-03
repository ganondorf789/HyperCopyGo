// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyPools is the golang structure of table proxy_pools for DAO operations like Where/Data.
type ProxyPools struct {
	g.Meta    `orm:"table:proxy_pools, do:true"`
	Id        any         // 主键ID
	Host      any         // 代理主机地址
	Port      any         // 代理端口
	Username  any         // 认证用户名
	Password  any         // 认证密码
	Status    any         // 状态（1=启用 0=禁用）
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
