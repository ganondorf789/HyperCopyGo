package admin

import (
	"context"

	"demo/api/admin/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	return service.Admin().Login(ctx, *req)
}
