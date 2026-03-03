package membership

import (
	"context"

	"demo/api/membership/v1"
	"demo/internal/service"
)

func (c *ControllerV1) MembershipCreate(ctx context.Context, req *v1.MembershipCreateReq) (res *v1.MembershipCreateRes, err error) {
	return service.Membership().Create(ctx, *req)
}
