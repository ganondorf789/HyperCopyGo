package service

import (
	"context"

	"demo/internal/model"
)

type IAdmin interface {
	Login(ctx context.Context, in model.AdminLoginInput) (*model.TokenOutput, error)
	Profile(ctx context.Context, adminId int64) (*model.AdminInfoOutput, error)
	UserList(ctx context.Context, in model.AdminUserListInput) (*model.AdminUserListOutput, error)
	UserSetStatus(ctx context.Context, in model.AdminUserStatusInput) error
	UserDelete(ctx context.Context, id int64) error
}

var localAdmin IAdmin

func Admin() IAdmin {
	return localAdmin
}

func RegisterAdmin(s IAdmin) {
	localAdmin = s
}
