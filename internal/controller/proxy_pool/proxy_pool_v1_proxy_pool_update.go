package proxy_pool

import (
	"context"

	"demo/api/proxy_pool/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ProxyPoolUpdate(ctx context.Context, req *v1.ProxyPoolUpdateReq) (res *v1.ProxyPoolUpdateRes, err error) {
	err = service.ProxyPool().Update(ctx, *req)
	return
}
