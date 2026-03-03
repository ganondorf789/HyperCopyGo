package model

import "github.com/gogf/gf/v2/os/gtime"

// NotificationCategorySummary 通知分类摘要（1级页面）
type NotificationCategorySummary struct {
	Category    string      `json:"category"`    // 通知类型 public/copy_trading/whale/track/market
	UnreadCount int         `json:"unreadCount"` // 未读消息数
	Latest      *NotificationPreview `json:"latest"`      // 最新一条消息预览，无消息时为 nil
}

// NotificationPreview 通知预览（用于摘要中的最新消息）
type NotificationPreview struct {
	Id        int64       `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	Level     int         `json:"level"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

// NotificationItem 通知列表项（2级页面）
type NotificationItem struct {
	Id        int64       `json:"id"`
	Category  string      `json:"category"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	RefId     int64       `json:"refId"`
	RefType   string      `json:"refType"`
	Level     int         `json:"level"`
	IsRead    bool        `json:"isRead"`
	CreatedAt *gtime.Time `json:"createdAt"`
}
