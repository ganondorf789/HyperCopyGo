package proxy_pool

import (
	"context"

	"demo/api/proxy_pool/v1"
	"demo/internal/service"
)

func (c *ControllerV1) ProxyPoolImportCSV(ctx context.Context, req *v1.ProxyPoolImportCSVReq) (res *v1.ProxyPoolImportCSVRes, err error) {
	return service.ProxyPool().ImportFromCSV(ctx, *req)
}
