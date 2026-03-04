package service

import (
	"context"

	v1 "demo/api/user/v1"
)

type IUser interface {
	SendVerifyCode(ctx context.Context, in v1.SendVerifyCodeReq) error
	Login(ctx context.Context, in v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	Profile(ctx context.Context, userId int64) (res *v1.UserProfileRes, err error)
}

var localUser IUser

func User() IUser {
	return localUser
}

func RegisterUser(s IUser) {
	localUser = s
}
