// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "demo/api/system_setting/v1"
)

type (
	ISystemSetting interface {
		Get(ctx context.Context) (res *v1.SystemSettingGetRes, err error)
		Update(ctx context.Context, in v1.SystemSettingUpdateReq) error
	}
)

var (
	localSystemSetting ISystemSetting
)

func SystemSetting() ISystemSetting {
	if localSystemSetting == nil {
		panic("implement not found for interface ISystemSetting, forgot register?")
	}
	return localSystemSetting
}

func RegisterSystemSetting(i ISystemSetting) {
	localSystemSetting = i
}
