package model

import "github.com/gogf/gf/v2/os/gtime"

// BaseProxyPool 代理池公共字段，供 API 请求/响应复用
type BaseProxyPool struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
}

// ProxyPoolItem 代理池列表项
type ProxyPoolItem struct {
	Id int64 `json:"id"`
	BaseProxyPool
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
