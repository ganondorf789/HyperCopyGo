package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 创建代理（需管理员权限）
type ProxyPoolCreateReq struct {
	g.Meta `path:"/proxy-pool" tags:"ProxyPool" method:"post" summary:"创建代理" login_required:"true" admin_required:"true"`
	model.BaseProxyPool
}
type ProxyPoolCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 更新代理（需管理员权限）
type ProxyPoolUpdateReq struct {
	g.Meta `path:"/proxy-pool/{id}" tags:"ProxyPool" method:"put" summary:"更新代理" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
	model.BaseProxyPool
	Status *int `json:"status"`
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

// 导入代理CSV（需管理员权限）
type ProxyPoolImportCSVReq struct {
	g.Meta `path:"/proxy-pool/import" tags:"ProxyPool" method:"post" summary:"导入代理CSV" login_required:"true" admin_required:"true"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required"`
}
type ProxyPoolImportCSVRes struct {
	g.Meta  `mime:"application/json"`
	Created int `json:"created"`
	Skipped int `json:"skipped"`
}

// 获取代理列表（需管理员权限）
type ProxyPoolListReq struct {
	g.Meta   `path:"/proxy-pool" tags:"ProxyPool" method:"get" summary:"代理列表" login_required:"true" admin_required:"true"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type ProxyPoolListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.ProxyPoolItem `json:"list"`
	Total  int                   `json:"total"`
	Page   int                   `json:"page"`
}
