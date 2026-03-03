// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package membership

import (
	"context"

	"demo/api/membership/v1"
)

type IMembershipV1 interface {
	MembershipCreate(ctx context.Context, req *v1.MembershipCreateReq) (res *v1.MembershipCreateRes, err error)
	MembershipUpdate(ctx context.Context, req *v1.MembershipUpdateReq) (res *v1.MembershipUpdateRes, err error)
	MembershipDelete(ctx context.Context, req *v1.MembershipDeleteReq) (res *v1.MembershipDeleteRes, err error)
	MembershipDetail(ctx context.Context, req *v1.MembershipDetailReq) (res *v1.MembershipDetailRes, err error)
	MembershipList(ctx context.Context, req *v1.MembershipListReq) (res *v1.MembershipListRes, err error)
}
