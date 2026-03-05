package proxy_pool

import (
	"context"

	"demo/api/proxy_pool/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ProxyPoolDelete(ctx context.Context, req *v1.ProxyPoolDeleteReq) (res *v1.ProxyPoolDeleteRes, err error) {
	err = service.ProxyPool().Delete(ctx, req.Id)
	return
}
