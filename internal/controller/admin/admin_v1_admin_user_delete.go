package admin

import (
	"context"

	"demo/api/admin/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteReq) (res *v1.AdminUserDeleteRes, err error) {
	err = service.Admin().UserDelete(ctx, req.Id)
	return
}
