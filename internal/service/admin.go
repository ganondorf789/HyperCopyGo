package service

import (
	"context"

	v1 "demo/api/admin/v1"
)

type IAdmin interface {
	Init(ctx context.Context, in v1.AdminInitReq) (res *v1.AdminInitRes, err error)
	Login(ctx context.Context, in v1.AdminLoginReq) (res *v1.AdminLoginRes, err error)
	Profile(ctx context.Context, adminId int64) (res *v1.AdminProfileRes, err error)
	UserList(ctx context.Context, in v1.AdminUserListReq) (res *v1.AdminUserListRes, err error)
	UserSearch(ctx context.Context, in v1.AdminUserSearchReq) (res *v1.AdminUserSearchRes, err error)
	UserSetStatus(ctx context.Context, in v1.AdminUserStatusReq) error
	UserDelete(ctx context.Context, id int64) error
}

var localAdmin IAdmin

func Admin() IAdmin {
	return localAdmin
}

func RegisterAdmin(s IAdmin) {
	localAdmin = s
}
