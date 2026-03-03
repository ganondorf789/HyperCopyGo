// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppVersion is the golang structure of table app_version for DAO operations like Where/Data.
type AppVersion struct {
	g.Meta         `orm:"table:app_version, do:true"`
	Id             any         // 主键ID
	Platform       any         // 平台 ios/android
	VersionName    any         // 版本号 如1.2.0
	VersionCode    any         // 版本编码 用于比较大小
	DownloadUrl    any         // 下载地址
	ChangeLog      any         // 更新日志
	ForceUpdate    any         // 是否强制更新 0:否 1:是
	MinVersionCode any         // 最低兼容版本编码 低于此版本强制更新
	Status         any         // 状态 1:已发布 0:未发布
	PublishedAt    *gtime.Time // 发布时间
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
