package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

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
