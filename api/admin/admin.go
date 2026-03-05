// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"demo/api/admin/v1"
)

type IAdminV1 interface {
	AdminInit(ctx context.Context, req *v1.AdminInitReq) (res *v1.AdminInitRes, err error)
	AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error)
	AdminProfile(ctx context.Context, req *v1.AdminProfileReq) (res *v1.AdminProfileRes, err error)
	AdminUserList(ctx context.Context, req *v1.AdminUserListReq) (res *v1.AdminUserListRes, err error)
	AdminUserSearch(ctx context.Context, req *v1.AdminUserSearchReq) (res *v1.AdminUserSearchRes, err error)
	AdminUserStatus(ctx context.Context, req *v1.AdminUserStatusReq) (res *v1.AdminUserStatusRes, err error)
	AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteReq) (res *v1.AdminUserDeleteRes, err error)
}
