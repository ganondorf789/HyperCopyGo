package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取系统设置（管理员）
type SystemSettingGetReq struct {
	g.Meta `path:"/system-setting" tags:"SystemSetting" method:"get" summary:"获取系统设置" login_required:"true" admin_required:"true"`
}
type SystemSettingGetRes struct {
	g.Meta `mime:"application/json"`
	model.SystemSettingItem
}

// 更新系统设置（管理员）
type SystemSettingUpdateReq struct {
	g.Meta                 `path:"/system-setting" tags:"SystemSetting" method:"put" summary:"更新系统设置" login_required:"true" admin_required:"true"`
	MarketMinutes          int64 `json:"marketMinutes"          v:"required|min:1#行情监控时间窗口必填|最小值为1分钟"`
	MarketNewPositionCount int64 `json:"marketNewPositionCount" v:"required|min:1#时间窗口内新仓位数量阈值必填|最小值为1"`
}
type SystemSettingUpdateRes struct {
	g.Meta `mime:"application/json"`
}
