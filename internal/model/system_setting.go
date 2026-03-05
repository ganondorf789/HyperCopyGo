package model

import "github.com/gogf/gf/v2/os/gtime"

// SystemSettingItem 系统设置
type SystemSettingItem struct {
	Id                     int64       `json:"id"`
	MarketMinutes          int64       `json:"marketMinutes"`
	MarketNewPositionCount int64       `json:"marketNewPositionCount"`
	CreatedAt              *gtime.Time `json:"createdAt"`
	UpdatedAt              *gtime.Time `json:"updatedAt"`
}
