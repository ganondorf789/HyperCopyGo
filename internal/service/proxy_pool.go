// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/proxy_pool/v1"
)

type (
	IProxyPool interface {
		Create(ctx context.Context, in v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error)
		Update(ctx context.Context, in v1.ProxyPoolUpdateReq) error
		Delete(ctx context.Context, id int64) error
		List(ctx context.Context, in v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error)
		ImportFromCSV(ctx context.Context, in v1.ProxyPoolImportCSVReq) (res *v1.ProxyPoolImportCSVRes, err error)
	}
)

var (
	localProxyPool IProxyPool
)

func ProxyPool() IProxyPool {
	if localProxyPool == nil {
		panic("implement not found for interface IProxyPool, forgot register?")
	}
	return localProxyPool
}

func RegisterProxyPool(i IProxyPool) {
	localProxyPool = i
}
