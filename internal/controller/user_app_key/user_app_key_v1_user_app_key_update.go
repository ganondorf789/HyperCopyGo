package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UserAppKeyUpdate(ctx context.Context, req *v1.UserAppKeyUpdateReq) (res *v1.UserAppKeyUpdateRes, err error) {
	return nil, service.UserAppKey().Update(ctx, *req)
}
