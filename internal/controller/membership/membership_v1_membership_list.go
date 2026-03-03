package membership

import (
	"context"

	"demo/api/membership/v1"
	"demo/internal/service"
)

func (c *ControllerV1) MembershipList(ctx context.Context, req *v1.MembershipListReq) (res *v1.MembershipListRes, err error) {
	return service.Membership().List(ctx, *req)
}
