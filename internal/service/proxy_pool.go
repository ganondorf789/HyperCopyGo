package service

import (
	"context"

	v1 "demo/api/proxy_pool/v1"
)

type IProxyPool interface {
	Create(ctx context.Context, in v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error)
	Update(ctx context.Context, in v1.ProxyPoolUpdateReq) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, in v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error)
}

var localProxyPool IProxyPool

func ProxyPool() IProxyPool {
	return localProxyPool
}

func RegisterProxyPool(s IProxyPool) {
	localProxyPool = s
}
