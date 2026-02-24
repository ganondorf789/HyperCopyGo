package service

import (
	"context"

	"demo/internal/model"
)

type IUser interface {
	Register(ctx context.Context, in model.UserRegisterInput) error
	Login(ctx context.Context, in model.UserLoginInput) (*model.TokenOutput, error)
	Profile(ctx context.Context, userId int64) (*model.UserInfoOutput, error)
}

var localUser IUser

func User() IUser {
	return localUser
}

func RegisterUser(s IUser) {
	localUser = s
}
