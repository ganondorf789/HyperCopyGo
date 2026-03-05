package admin

import (
	"context"

	"demo/api/admin/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AdminUserStatus(ctx context.Context, req *v1.AdminUserStatusReq) (res *v1.AdminUserStatusRes, err error) {
	err = service.Admin().UserSetStatus(ctx, *req)
	return
}
