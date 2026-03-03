// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVersion is the golang structure for table app_version.
type AppVersion struct {
	Id             int64       `json:"id"             orm:"id"               description:"主键ID"`               // 主键ID
	Platform       string      `json:"platform"       orm:"platform"         description:"平台 ios/android"`     // 平台 ios/android
	VersionName    string      `json:"versionName"    orm:"version_name"     description:"版本号 如1.2.0"`         // 版本号 如1.2.0
	VersionCode    int64       `json:"versionCode"    orm:"version_code"     description:"版本编码 用于比较大小"`        // 版本编码 用于比较大小
	DownloadUrl    string      `json:"downloadUrl"    orm:"download_url"     description:"下载地址"`               // 下载地址
	ChangeLog      string      `json:"changeLog"      orm:"change_log"       description:"更新日志"`               // 更新日志
	ForceUpdate    int         `json:"forceUpdate"    orm:"force_update"     description:"是否强制更新 0:否 1:是"`     // 是否强制更新 0:否 1:是
	MinVersionCode int64       `json:"minVersionCode" orm:"min_version_code" description:"最低兼容版本编码 低于此版本强制更新"` // 最低兼容版本编码 低于此版本强制更新
	Status         int         `json:"status"         orm:"status"           description:"状态 1:已发布 0:未发布"`     // 状态 1:已发布 0:未发布
	PublishedAt    *gtime.Time `json:"publishedAt"    orm:"published_at"     description:"发布时间"`               // 发布时间
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`               // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`               // 更新时间
}
