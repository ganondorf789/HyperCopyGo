// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/membership/v1"
)

type (
	IMembership interface {
		Create(ctx context.Context, in v1.MembershipCreateReq) (res *v1.MembershipCreateRes, err error)
		Update(ctx context.Context, in v1.MembershipUpdateReq) error
		Delete(ctx context.Context, id int64) error
		Detail(ctx context.Context, id int64) (res *v1.MembershipDetailRes, err error)
		List(ctx context.Context, in v1.MembershipListReq) (res *v1.MembershipListRes, err error)
	}
)

var (
	localMembership IMembership
)

func Membership() IMembership {
	if localMembership == nil {
		panic("implement not found for interface IMembership, forgot register?")
	}
	return localMembership
}

func RegisterMembership(i IMembership) {
	localMembership = i
}
