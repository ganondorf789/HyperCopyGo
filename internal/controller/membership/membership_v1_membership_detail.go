package membership

import (
	"context"

	"demo/api/membership/v1"
	"demo/internal/service"
)

func (c *ControllerV1) MembershipDetail(ctx context.Context, req *v1.MembershipDetailReq) (res *v1.MembershipDetailRes, err error) {
	return service.Membership().Detail(ctx, req.Id)
}
