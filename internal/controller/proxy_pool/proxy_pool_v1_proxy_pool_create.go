package proxy_pool

import (
	"context"

	"demo/api/proxy_pool/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ProxyPoolCreate(ctx context.Context, req *v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error) {
	return service.ProxyPool().Create(ctx, *req)
}
