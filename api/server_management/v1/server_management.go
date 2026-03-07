package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ServerManagementCreateReq struct {
	g.Meta   `path:"/server-management" tags:"ServerManagement" method:"post" summary:"创建服务器" login_required:"true" admin_required:"true"`
	Ip       string `json:"ip" v:"required#请输入服务器IP"`
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}
type ServerManagementCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

type ServerManagementUpdateReq struct {
	g.Meta   `path:"/server-management/{id}" tags:"ServerManagement" method:"put" summary:"更新服务器" login_required:"true" admin_required:"true"`
	Id       int64  `json:"id" in:"path" v:"required"`
	Ip       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type ServerManagementUpdateRes struct {
	g.Meta `mime:"application/json"`
}

type ServerManagementDeleteReq struct {
	g.Meta `path:"/server-management/{id}" tags:"ServerManagement" method:"delete" summary:"删除服务器" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type ServerManagementDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type ServerManagementDetailReq struct {
	g.Meta `path:"/server-management/{id}" tags:"ServerManagement" method:"get" summary:"服务器详情" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type ServerManagementDetailRes struct {
	g.Meta `mime:"application/json"`
	ServerManagementItem
}

type ServerManagementListReq struct {
	g.Meta   `path:"/server-management" tags:"ServerManagement" method:"get" summary:"服务器列表" login_required:"true" admin_required:"true"`
	Page     int `json:"page" in:"query" d:"1"`
	PageSize int `json:"pageSize" in:"query" d:"20" v:"max:100#每页最多100条"`
}
type ServerManagementListRes struct {
	g.Meta `mime:"application/json"`
	List   []ServerManagementItem `json:"list"`
	Total  int                    `json:"total"`
	Page   int                    `json:"page"`
}

type ServerManagementItem struct {
	Id        int64       `json:"id"`
	Ip        string      `json:"ip"`
	Username  string      `json:"username"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
