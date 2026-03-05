// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package proxy_pool

import (
	"context"

	"demo/api/proxy_pool/v1"
)

type IProxyPoolV1 interface {
	ProxyPoolCreate(ctx context.Context, req *v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error)
	ProxyPoolUpdate(ctx context.Context, req *v1.ProxyPoolUpdateReq) (res *v1.ProxyPoolUpdateRes, err error)
	ProxyPoolDelete(ctx context.Context, req *v1.ProxyPoolDeleteReq) (res *v1.ProxyPoolDeleteRes, err error)
	ProxyPoolImportCSV(ctx context.Context, req *v1.ProxyPoolImportCSVReq) (res *v1.ProxyPoolImportCSVRes, err error)
	ProxyPoolList(ctx context.Context, req *v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error)
}
