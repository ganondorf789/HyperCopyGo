package v1

import (
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建代理（需管理员权限）
type ProxyPoolCreateReq struct {
	g.Meta   `path:"/proxy-pool" tags:"ProxyPool" method:"post" summary:"创建代理" login_required:"true" admin_required:"true"`
	Host     string `json:"host"     v:"required#请输入代理主机地址"`
	Port     int    `json:"port"     v:"required#请输入代理端口"`
	Username string `json:"username" v:"required#请输入认证用户名"`
	Password string `json:"password" v:"required#请输入认证密码"`
	Remark   string `json:"remark"`
}
type ProxyPoolCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 更新代理（需管理员权限）
type ProxyPoolUpdateReq struct {
	g.Meta   `path:"/proxy-pool/{id}" tags:"ProxyPool" method:"put" summary:"更新代理" login_required:"true" admin_required:"true"`
	Id       int64  `json:"id"       in:"path" v:"required"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   *int   `json:"status"`
	Remark   string `json:"remark"`
}
type ProxyPoolUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除代理（需管理员权限）
type ProxyPoolDeleteReq struct {
	g.Meta `path:"/proxy-pool/{id}" tags:"ProxyPool" method:"delete" summary:"删除代理" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type ProxyPoolDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 获取代理列表（需管理员权限）
type ProxyPoolListReq struct {
	g.Meta   `path:"/proxy-pool" tags:"ProxyPool" method:"get" summary:"代理列表" login_required:"true" admin_required:"true"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type ProxyPoolListRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.ProxyPool `json:"list"`
	Total  int                `json:"total"`
	Page   int                `json:"page"`
}
