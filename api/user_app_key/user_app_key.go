// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
)

type IUserAppKeyV1 interface {
	UserAppKeyCreate(ctx context.Context, req *v1.UserAppKeyCreateReq) (res *v1.UserAppKeyCreateRes, err error)
	UserAppKeyRefreshSecret(ctx context.Context, req *v1.UserAppKeyRefreshSecretReq) (res *v1.UserAppKeyRefreshSecretRes, err error)
	UserAppKeyUpdate(ctx context.Context, req *v1.UserAppKeyUpdateReq) (res *v1.UserAppKeyUpdateRes, err error)
	UserAppKeyDelete(ctx context.Context, req *v1.UserAppKeyDeleteReq) (res *v1.UserAppKeyDeleteRes, err error)
	UserAppKeyDetail(ctx context.Context, req *v1.UserAppKeyDetailReq) (res *v1.UserAppKeyDetailRes, err error)
	UserAppKeyList(ctx context.Context, req *v1.UserAppKeyListReq) (res *v1.UserAppKeyListRes, err error)
}
