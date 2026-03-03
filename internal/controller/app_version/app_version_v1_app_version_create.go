package app_version

import (
	"context"

	"demo/api/app_version/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AppVersionCreate(ctx context.Context, req *v1.AppVersionCreateReq) (res *v1.AppVersionCreateRes, err error) {
	return service.AppVersion().Create(ctx, *req)
}
