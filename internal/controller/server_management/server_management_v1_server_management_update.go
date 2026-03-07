package server_management

import (
	"context"

	"demo/api/server_management/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ServerManagementUpdate(ctx context.Context, req *v1.ServerManagementUpdateReq) (res *v1.ServerManagementUpdateRes, err error) {
	if err = service.ServerManagement().Update(ctx, *req); err != nil {
		return nil, err
	}
	return &v1.ServerManagementUpdateRes{}, nil
}
