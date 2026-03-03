package app_version

import (
	"context"

	"demo/api/app_version/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AppVersionUpdate(ctx context.Context, req *v1.AppVersionUpdateReq) (res *v1.AppVersionUpdateRes, err error) {
	return nil, service.AppVersion().Update(ctx, *req)
}
