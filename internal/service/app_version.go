// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/app_version/v1"
)

type (
	IAppVersion interface {
		Create(ctx context.Context, in v1.AppVersionCreateReq) (res *v1.AppVersionCreateRes, err error)
		Update(ctx context.Context, in v1.AppVersionUpdateReq) error
		Delete(ctx context.Context, id int64) error
		List(ctx context.Context, in v1.AppVersionListReq) (res *v1.AppVersionListRes, err error)
		Check(ctx context.Context, in v1.AppVersionCheckReq) (res *v1.AppVersionCheckRes, err error)
	}
)

var (
	localAppVersion IAppVersion
)

func AppVersion() IAppVersion {
	if localAppVersion == nil {
		panic("implement not found for interface IAppVersion, forgot register?")
	}
	return localAppVersion
}

func RegisterAppVersion(i IAppVersion) {
	localAppVersion = i
}
