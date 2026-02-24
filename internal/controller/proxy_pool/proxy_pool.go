package proxy_pool

import (
	"context"

	v1 "demo/api/proxy_pool/v1"
	"demo/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) ProxyPoolCreate(ctx context.Context, req *v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error) {
	return service.ProxyPool().Create(ctx, *req)
}

func (c *Controller) ProxyPoolUpdate(ctx context.Context, req *v1.ProxyPoolUpdateReq) (res *v1.ProxyPoolUpdateRes, err error) {
	err = service.ProxyPool().Update(ctx, *req)
	return
}

func (c *Controller) ProxyPoolDelete(ctx context.Context, req *v1.ProxyPoolDeleteReq) (res *v1.ProxyPoolDeleteRes, err error) {
	err = service.ProxyPool().Delete(ctx, req.Id)
	return
}

func (c *Controller) ProxyPoolList(ctx context.Context, req *v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error) {
	return service.ProxyPool().List(ctx, *req)
}
