package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 创建APP版本（管理员）
type AppVersionCreateReq struct {
	g.Meta `path:"/app-version" tags:"AppVersion" method:"post" summary:"创建APP版本" login_required:"true" admin_required:"true"`
	model.BaseAppVersion
	Status int `json:"status" d:"0"`
}
type AppVersionCreateRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id"`
}

// 更新APP版本（管理员）
type AppVersionUpdateReq struct {
	g.Meta `path:"/app-version/{id}" tags:"AppVersion" method:"put" summary:"更新APP版本" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
	model.BaseAppVersion
	Status int `json:"status"`
}
type AppVersionUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 删除APP版本（管理员）
type AppVersionDeleteReq struct {
	g.Meta `path:"/app-version/{id}" tags:"AppVersion" method:"delete" summary:"删除APP版本" login_required:"true" admin_required:"true"`
	Id     int64 `json:"id" in:"path" v:"required"`
}
type AppVersionDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// APP版本列表（管理员）
type AppVersionListReq struct {
	g.Meta   `path:"/app-version" tags:"AppVersion" method:"get" summary:"APP版本列表" login_required:"true" admin_required:"true"`
	Platform string `json:"platform"`
	Status   int    `json:"status" d:"-1"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20" v:"max:100#每页最多100条"`
}
type AppVersionListRes struct {
	g.Meta `mime:"application/json"`
	List   []AppVersionItem `json:"list"`
	Total  int              `json:"total"`
	Page   int              `json:"page"`
}

// 检查APP更新（免鉴权）
type AppVersionCheckReq struct {
	g.Meta      `path:"/app-version/check" tags:"AppVersion" method:"get" summary:"检查APP更新"`
	Platform    string `json:"platform" v:"required|in:ios,android#请选择平台#平台只能是ios或android"`
	VersionCode int64  `json:"versionCode" v:"required#请输入当前版本编码"`
}
type AppVersionCheckRes struct {
	g.Meta      `mime:"application/json"`
	HasUpdate   bool   `json:"hasUpdate"`
	ForceUpdate bool   `json:"forceUpdate"`
	VersionName string `json:"versionName,omitempty"`
	VersionCode int64  `json:"versionCode,omitempty"`
	DownloadUrl string `json:"downloadUrl,omitempty"`
	ChangeLog   string `json:"changeLog,omitempty"`
}

// 列表项
type AppVersionItem struct {
	Id int64 `json:"id"`
	model.BaseAppVersion
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}
