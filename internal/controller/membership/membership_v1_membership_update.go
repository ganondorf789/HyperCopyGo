package membership

import (
	"context"

	"demo/api/membership/v1"
	"demo/internal/service"
)

func (c *ControllerV1) MembershipUpdate(ctx context.Context, req *v1.MembershipUpdateReq) (res *v1.MembershipUpdateRes, err error) {
	return nil, service.Membership().Update(ctx, *req)
}
