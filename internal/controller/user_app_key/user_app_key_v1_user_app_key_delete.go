package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UserAppKeyDelete(ctx context.Context, req *v1.UserAppKeyDeleteReq) (res *v1.UserAppKeyDeleteRes, err error) {
	return nil, service.UserAppKey().Delete(ctx, req.Id)
}
