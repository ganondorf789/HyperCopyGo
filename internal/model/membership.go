package model

import "github.com/gogf/gf/v2/os/gtime"

// MembershipItem 会员列表项
type MembershipItem struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	Level     int         `json:"level"`
	StartAt   *gtime.Time `json:"startAt"`
	ExpireAt  *gtime.Time `json:"expireAt"`
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
