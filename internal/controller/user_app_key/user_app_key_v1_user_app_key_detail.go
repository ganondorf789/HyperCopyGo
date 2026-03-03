package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UserAppKeyDetail(ctx context.Context, req *v1.UserAppKeyDetailReq) (res *v1.UserAppKeyDetailRes, err error) {
	return service.UserAppKey().Detail(ctx, req.Id)
}
