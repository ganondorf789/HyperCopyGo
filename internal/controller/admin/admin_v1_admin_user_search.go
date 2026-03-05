package admin

import (
	"context"

	"demo/api/admin/v1"
	"demo/internal/service"
)

func (c *ControllerV1) AdminUserSearch(ctx context.Context, req *v1.AdminUserSearchReq) (res *v1.AdminUserSearchRes, err error) {
	return service.Admin().UserSearch(ctx, *req)
}
