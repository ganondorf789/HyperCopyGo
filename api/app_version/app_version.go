// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app_version

import (
	"context"

	"demo/api/app_version/v1"
)

type IAppVersionV1 interface {
	AppVersionCreate(ctx context.Context, req *v1.AppVersionCreateReq) (res *v1.AppVersionCreateRes, err error)
	AppVersionUpdate(ctx context.Context, req *v1.AppVersionUpdateReq) (res *v1.AppVersionUpdateRes, err error)
	AppVersionDelete(ctx context.Context, req *v1.AppVersionDeleteReq) (res *v1.AppVersionDeleteRes, err error)
	AppVersionList(ctx context.Context, req *v1.AppVersionListReq) (res *v1.AppVersionListRes, err error)
	AppVersionCheck(ctx context.Context, req *v1.AppVersionCheckReq) (res *v1.AppVersionCheckRes, err error)
}
