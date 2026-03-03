// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"demo/api/user/v1"
)

type IUserV1 interface {
	UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error)
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error)
}
