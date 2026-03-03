package model

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
