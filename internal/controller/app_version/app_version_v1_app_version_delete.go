package app_version

import (
	"context"

	"demo/api/app_version/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AppVersionDelete(ctx context.Context, req *v1.AppVersionDeleteReq) (res *v1.AppVersionDeleteRes, err error) {
	return nil, service.AppVersion().Delete(ctx, req.Id)
}
