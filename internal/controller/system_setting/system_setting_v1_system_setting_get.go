package system_setting

import (
	"context"

	v1 "demo/api/system_setting/v1"
	"demo/internal/service"
)

func (c *ControllerV1) SystemSettingGet(ctx context.Context, req *v1.SystemSettingGetReq) (res *v1.SystemSettingGetRes, err error) {
	return service.SystemSetting().Get(ctx)
}
