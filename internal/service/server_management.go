// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/server_management/v1"
)

type (
	IServerManagement interface {
		Create(ctx context.Context, in v1.ServerManagementCreateReq) (res *v1.ServerManagementCreateRes, err error)
		Update(ctx context.Context, in v1.ServerManagementUpdateReq) error
		Delete(ctx context.Context, id int64) error
		Detail(ctx context.Context, id int64) (res *v1.ServerManagementDetailRes, err error)
		List(ctx context.Context, in v1.ServerManagementListReq) (res *v1.ServerManagementListRes, err error)
	}
)

var (
	localServerManagement IServerManagement
)

func ServerManagement() IServerManagement {
	if localServerManagement == nil {
		panic("implement not found for interface IServerManagement, forgot register?")
	}
	return localServerManagement
}

func RegisterServerManagement(i IServerManagement) {
	localServerManagement = i
}
