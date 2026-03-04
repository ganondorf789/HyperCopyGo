package user_app_key

import (
	"context"

	"demo/api/user_app_key/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UserAppKeyRefreshSecret(ctx context.Context, req *v1.UserAppKeyRefreshSecretReq) (res *v1.UserAppKeyRefreshSecretRes, err error) {
	return service.UserAppKey().RefreshSecret(ctx, *req)
}
