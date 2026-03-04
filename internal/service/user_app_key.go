// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/user_app_key/v1"
)

type (
	IUserAppKey interface {
		Create(ctx context.Context, in v1.UserAppKeyCreateReq) (res *v1.UserAppKeyCreateRes, err error)
		RefreshSecret(ctx context.Context, in v1.UserAppKeyRefreshSecretReq) (res *v1.UserAppKeyRefreshSecretRes, err error)
		Update(ctx context.Context, in v1.UserAppKeyUpdateReq) error
		Delete(ctx context.Context, id int64) error
		Detail(ctx context.Context, id int64) (res *v1.UserAppKeyDetailRes, err error)
		List(ctx context.Context, in v1.UserAppKeyListReq) (res *v1.UserAppKeyListRes, err error)
	}
)

var (
	localUserAppKey IUserAppKey
)

func UserAppKey() IUserAppKey {
	if localUserAppKey == nil {
		panic("implement not found for interface IUserAppKey, forgot register?")
	}
	return localUserAppKey
}

func RegisterUserAppKey(i IUserAppKey) {
	localUserAppKey = i
}
