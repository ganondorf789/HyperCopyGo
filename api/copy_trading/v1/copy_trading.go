package v1

import (
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 创建跟单配置（需登录）
type CopyTradingCreateReq struct {
	g.Meta `path:"/copy-trading" tags:"CopyTrading" method:"post" summary:"创建跟单配置" login_required:"true"`
	entity.BaseCopyTrading
}
type CopyTradingCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 更新跟单配置（需登录）
type CopyTradingUpdateReq struct {
	g.Meta `path:"/copy-trading/{id}" tags:"CopyTrading" method:"put" summary:"更新跟单配置" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
	entity.BaseCopyTrading
	Status int `json:"status"`
}
type CopyTradingUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除跟单配置（需登录）
type CopyTradingDeleteReq struct {
	g.Meta `path:"/copy-trading/{id}" tags:"CopyTrading" method:"delete" summary:"删除跟单配置" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type CopyTradingDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 获取跟单配置详情（需登录）
type CopyTradingDetailReq struct {
	g.Meta `path:"/copy-trading/{id}" tags:"CopyTrading" method:"get" summary:"跟单配置详情" login_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type CopyTradingDetailRes struct {
	g.Meta `mime:"application/json"`
	CopyTradingItem
}

// 获取跟单配置列表（需登录）
type CopyTradingListReq struct {
	g.Meta   `path:"/copy-trading" tags:"CopyTrading" method:"get" summary:"跟单配置列表" login_required:"true"`
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
	Status   int `json:"status" d:"-1"`
}
type CopyTradingListRes struct {
	g.Meta `mime:"application/json"`
	List   []CopyTradingItem `json:"list"`
	Total  int               `json:"total"`
	Page   int               `json:"page"`
}

// 列表项
type CopyTradingItem struct {
	Id int64 `json:"id"`
	entity.BaseCopyTrading
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
