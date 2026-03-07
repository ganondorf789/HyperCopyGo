package server_management

import (
	"context"

	"demo/api/server_management/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ServerManagementDelete(ctx context.Context, req *v1.ServerManagementDeleteReq) (res *v1.ServerManagementDeleteRes, err error) {
	if err = service.ServerManagement().Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.ServerManagementDeleteRes{}, nil
}
