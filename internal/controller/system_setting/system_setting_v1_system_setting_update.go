package system_setting

import (
	"context"

	v1 "demo/api/system_setting/v1"
	"demo/internal/service"
)

func (c *ControllerV1) SystemSettingUpdate(ctx context.Context, req *v1.SystemSettingUpdateReq) (res *v1.SystemSettingUpdateRes, err error) {
	return nil, service.SystemSetting().Update(ctx, *req)
}
