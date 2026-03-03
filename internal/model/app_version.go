package model

import "github.com/gogf/gf/v2/os/gtime"

// BaseAppVersion APP版本公共字段，供 API 请求/响应复用
type BaseAppVersion struct {
	Platform       string `json:"platform"`
	VersionName    string `json:"versionName"`
	VersionCode    int64  `json:"versionCode"`
	DownloadUrl    string `json:"downloadUrl"`
	ChangeLog      string `json:"changeLog"`
	ForceUpdate    int    `json:"forceUpdate"`
	MinVersionCode int64  `json:"minVersionCode"`
}

// AppVersionItem APP版本列表项
type AppVersionItem struct {
	Id int64 `json:"id"`
	BaseAppVersion
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
