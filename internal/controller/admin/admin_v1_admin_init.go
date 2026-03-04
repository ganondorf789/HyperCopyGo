package admin

import (
	"context"

	"demo/api/admin/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AdminInit(ctx context.Context, req *v1.AdminInitReq) (res *v1.AdminInitRes, err error) {
	return service.Admin().Init(ctx, *req)
}
