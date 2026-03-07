package server_management

import (
	"context"

	"demo/api/server_management/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ServerManagementCreate(ctx context.Context, req *v1.ServerManagementCreateReq) (res *v1.ServerManagementCreateRes, err error) {
	return service.ServerManagement().Create(ctx, *req)
}
