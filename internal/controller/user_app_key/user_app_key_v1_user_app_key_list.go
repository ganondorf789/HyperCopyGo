package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UserAppKeyList(ctx context.Context, req *v1.UserAppKeyListReq) (res *v1.UserAppKeyListRes, err error) {
	return service.UserAppKey().List(ctx, *req)
}
