package server_management

import (
	"context"

	"demo/api/server_management/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ServerManagementList(ctx context.Context, req *v1.ServerManagementListReq) (res *v1.ServerManagementListRes, err error) {
	return service.ServerManagement().List(ctx, *req)
}
