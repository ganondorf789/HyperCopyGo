package admin

import (
	"context"

	v1 "demo/api/admin/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	return service.Admin().Login(ctx, *req)
}

func (c *Controller) AdminProfile(ctx context.Context, req *v1.AdminProfileReq) (res *v1.AdminProfileRes, err error) {
	adminId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.Admin().Profile(ctx, adminId)
}

func (c *Controller) AdminUserList(ctx context.Context, req *v1.AdminUserListReq) (res *v1.AdminUserListRes, err error) {
	return service.Admin().UserList(ctx, *req)
}

func (c *Controller) AdminUserStatus(ctx context.Context, req *v1.AdminUserStatusReq) (res *v1.AdminUserStatusRes, err error) {
	err = service.Admin().UserSetStatus(ctx, *req)
	return
}

func (c *Controller) AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteReq) (res *v1.AdminUserDeleteRes, err error) {
	err = service.Admin().UserDelete(ctx, req.Id)
	return
}
