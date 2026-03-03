package app_version

import (
	"context"

	"demo/api/app_version/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AppVersionList(ctx context.Context, req *v1.AppVersionListReq) (res *v1.AppVersionListRes, err error) {
	return service.AppVersion().List(ctx, *req)
}
