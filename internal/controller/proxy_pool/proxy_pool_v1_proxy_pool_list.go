package proxy_pool

import (
	"context"

	"demo/api/proxy_pool/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ProxyPoolList(ctx context.Context, req *v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error) {
	return service.ProxyPool().List(ctx, *req)
}
