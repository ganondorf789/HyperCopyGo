package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 创建用户AppKey（管理员）
type UserAppKeyCreateReq struct {
	g.Meta   `path:"/user-app-key" tags:"UserAppKey" method:"post" summary:"创建用户AppKey" login_required:"true" admin_required:"true"`
	UserId   int64       `json:"userId" v:"required#请输入用户ID"`
	Remark   string      `json:"remark"`
	ExpireAt *gtime.Time `json:"expireAt"`
	Status   int         `json:"status" d:"1"`
}
type UserAppKeyCreateRes struct {
	g.Meta    `mime:"application/json"`
	Id        int64  `json:"id"`
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

// 更新用户AppKey（管理员）
type UserAppKeyUpdateReq struct {
	g.Meta   `path:"/user-app-key/{id}" tags:"UserAppKey" method:"put" summary:"更新用户AppKey" login_required:"true" admin_required:"true"`
	Id       int64       `json:"id" in:"path" v:"required"`
	Remark   string      `json:"remark"`
	ExpireAt *gtime.Time `json:"expireAt"`
	Status   int         `json:"status"`
}
type UserAppKeyUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除用户AppKey（管理员）
type UserAppKeyDeleteReq struct {
	g.Meta `path:"/user-app-key/{id}" tags:"UserAppKey" method:"delete" summary:"删除用户AppKey" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type UserAppKeyDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 用户AppKey详情（管理员）
type UserAppKeyDetailReq struct {
	g.Meta `path:"/user-app-key/{id}" tags:"UserAppKey" method:"get" summary:"用户AppKey详情" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type UserAppKeyDetailRes struct {
	g.Meta `mime:"application/json"`
	model.UserAppKeyItem
}

// 用户AppKey列表（管理员）
type UserAppKeyListReq struct {
	g.Meta   `path:"/user-app-key" tags:"UserAppKey" method:"get" summary:"用户AppKey列表" login_required:"true" admin_required:"true"`
	UserId   int64 `json:"userId"`
	Status   int   `json:"status" d:"-1"`
	Page     int   `json:"page" d:"1"`
	PageSize int   `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type UserAppKeyListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.UserAppKeyItem `json:"list"`
	Total  int                    `json:"total"`
	Page   int                    `json:"page"`
}