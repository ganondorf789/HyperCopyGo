package membership

import (
	"context"

	"demo/api/membership/v1"
	"demo/internal/service"
)

func (c *ControllerV1) MembershipDelete(ctx context.Context, req *v1.MembershipDeleteReq) (res *v1.MembershipDeleteRes, err error) {
	return nil, service.Membership().Delete(ctx, req.Id)
}
