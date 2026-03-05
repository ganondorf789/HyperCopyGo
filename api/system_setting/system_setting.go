// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system_setting

import (
	"context"

	"demo/api/system_setting/v1"
)

type ISystemSettingV1 interface {
	SystemSettingGet(ctx context.Context, req *v1.SystemSettingGetReq) (res *v1.SystemSettingGetRes, err error)
	SystemSettingUpdate(ctx context.Context, req *v1.SystemSettingUpdateReq) (res *v1.SystemSettingUpdateRes, err error)
}
