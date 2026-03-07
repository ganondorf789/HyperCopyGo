// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package server_management

import (
	"context"

	"demo/api/server_management/v1"
)

type IServerManagementV1 interface {
	ServerManagementCreate(ctx context.Context, req *v1.ServerManagementCreateReq) (res *v1.ServerManagementCreateRes, err error)
	ServerManagementUpdate(ctx context.Context, req *v1.ServerManagementUpdateReq) (res *v1.ServerManagementUpdateRes, err error)
	ServerManagementDelete(ctx context.Context, req *v1.ServerManagementDeleteReq) (res *v1.ServerManagementDeleteRes, err error)
	ServerManagementDetail(ctx context.Context, req *v1.ServerManagementDetailReq) (res *v1.ServerManagementDetailRes, err error)
	ServerManagementList(ctx context.Context, req *v1.ServerManagementListReq) (res *v1.ServerManagementListRes, err error)
}
