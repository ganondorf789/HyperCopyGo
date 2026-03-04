// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"demo/api/user/v1"
)

type IUserV1 interface {
	SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeReq) (res *v1.SendVerifyCodeRes, err error)
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error)
}
