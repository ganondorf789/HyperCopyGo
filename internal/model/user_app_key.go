package model

import "github.com/gogf/gf/v2/os/gtime"

// UserAppKeyItem 用户AppKey列表项
type UserAppKeyItem struct {
	Id        int64       `json:"id"`
	UserId    int64       `json:"userId"`
	AppId     string      `json:"appId"`
	AppSecret string      `json:"appSecret"`
	Remark    string      `json:"remark"`
	ExpireAt  *gtime.Time `json:"expireAt"`
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
