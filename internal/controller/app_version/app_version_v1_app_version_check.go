package app_version

import (
	"context"

	"demo/api/app_version/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AppVersionCheck(ctx context.Context, req *v1.AppVersionCheckReq) (res *v1.AppVersionCheckRes, err error) {
	return service.AppVersion().Check(ctx, *req)
}
