package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 发送公共通知（管理员）
type NotificationSendReq struct {
	g.Meta  `path:"/notification/send" tags:"Notification" method:"post" summary:"发送公共通知" login_required:"true" admin_required:"true"`
	Category string `json:"category" v:"required|in:public,copy_trading,whale,track,market#请选择通知类型#通知类型不合法"`
	Title    string `json:"title" v:"required#请输入通知标题"`
	Content  string `json:"content" v:"required#请输入通知内容"`
	Level    int    `json:"level" d:"0"`
	RefId    int64  `json:"refId" d:"0"`
	RefType  string `json:"refType"`
}
type NotificationSendRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 编辑通知（管理员）
type NotificationUpdateReq struct {
	g.Meta  `path:"/notification/{id}" tags:"Notification" method:"put" summary:"编辑通知" login_required:"true" admin_required:"true"`
	Id      int64  `json:"id" in:"path" v:"required"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Level    int    `json:"level"`
	RefId    int64  `json:"refId"`
	RefType  string `json:"refType"`
	Status   int    `json:"status"`
}
type NotificationUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除通知（管理员）
type NotificationDeleteReq struct {
	g.Meta `path:"/notification/{id}" tags:"Notification" method:"delete" summary:"删除通知" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type NotificationDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// 通知管理列表（管理员，查看所有通知）
type NotificationAdminListReq struct {
	g.Meta   `path:"/notification/admin" tags:"Notification" method:"get" summary:"通知管理列表" login_required:"true" admin_required:"true"`
	Category string `json:"category"`
	Status   int    `json:"status" d:"-1"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type NotificationAdminListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.NotificationAdminItem `json:"list"`
	Total  int                           `json:"total"`
	Page   int                           `json:"page"`
}

// 通知分类摘要（1级页面：各分类未读数 + 最新消息预览）
type NotificationSummaryReq struct {
	g.Meta `path:"/notification/summary" tags:"Notification" method:"get" summary:"通知分类摘要" login_required:"true"`
}
type NotificationSummaryRes struct {
	g.Meta     `mime:"application/json"`
	Categories []model.NotificationCategorySummary `json:"categories"`
}

// 分类通知列表（2级页面）
type NotificationListReq struct {
	g.Meta   `path:"/notification" tags:"Notification" method:"get" summary:"通知列表" login_required:"true"`
	Category string `json:"category" v:"required|in:public,copy_trading,whale,track,market#请选择通知类型#通知类型不合法"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type NotificationListRes struct {
	g.Meta `mime:"application/json"`
	List   []model.NotificationItem `json:"list"`
	Total  int                      `json:"total"`
	Page   int                      `json:"page"`
}

// 标记已读（单条/批量）
type NotificationReadReq struct {
	g.Meta `path:"/notification/read" tags:"Notification" method:"put" summary:"标记通知已读" login_required:"true"`
	Ids    []int64 `json:"ids" v:"required#请传入通知ID列表"`
}
type NotificationReadRes struct {
	g.Meta `mime:"application/json"`
}

// 全部已读（按分类）
type NotificationReadAllReq struct {
	g.Meta   `path:"/notification/read-all" tags:"Notification" method:"put" summary:"全部标记已读" login_required:"true"`
	Category string `json:"category" v:"required|in:public,copy_trading,whale,track,market#请选择通知类型#通知类型不合法"`
}
type NotificationReadAllRes struct {
	g.Meta `mime:"application/json"`
}
