package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UserAppKeyCreate(ctx context.Context, req *v1.UserAppKeyCreateReq) (res *v1.UserAppKeyCreateRes, err error) {
	return service.UserAppKey().Create(ctx, *req)
}
