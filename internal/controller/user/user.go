package user

import (
	"context"

	v1 "demo/api/user/v1"
	"demo/internal/consts"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeReq) (res *v1.SendVerifyCodeRes, err error) {
	err = service.User().SendVerifyCode(ctx, *req)
	return
}

func (c *Controller) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	return service.User().Login(ctx, *req)
}

func (c *Controller) UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	return service.User().Profile(ctx, userId)
}
