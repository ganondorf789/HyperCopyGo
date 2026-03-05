package user

import (
	"context"

	"demo/api/user/v1"
	"demo/internal/service"
)

func (c *ControllerV1) SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeReq) (res *v1.SendVerifyCodeRes, err error) {
	err = service.User().SendVerifyCode(ctx, *req)
	return
}
