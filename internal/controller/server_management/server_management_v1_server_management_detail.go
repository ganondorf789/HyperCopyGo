package server_management

import (
	"context"

	"demo/api/server_management/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ServerManagementDetail(ctx context.Context, req *v1.ServerManagementDetailReq) (res *v1.ServerManagementDetailRes, err error) {
	return service.ServerManagement().Detail(ctx, req.Id)
}
