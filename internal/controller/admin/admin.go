package admin

import (
	"context"

	v1 "demo/api/admin/v1"
	"demo/internal/consts"
	"demo/internal/model"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	out, err := service.Admin().Login(ctx, model.AdminLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AdminLoginRes{
		Token:  out.Token,
		Expire: out.Expire,
	}, nil
}

func (c *Controller) AdminProfile(ctx context.Context, req *v1.AdminProfileReq) (res *v1.AdminProfileRes, err error) {
	adminId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	out, err := service.Admin().Profile(ctx, adminId)
	if err != nil {
		return nil, err
	}
	return &v1.AdminProfileRes{
		Id:       out.Id,
		Username: out.Username,
		Realname: out.Realname,
		Role:     out.Role,
	}, nil
}

func (c *Controller) AdminUserList(ctx context.Context, req *v1.AdminUserListReq) (res *v1.AdminUserListRes, err error) {
	out, err := service.Admin().UserList(ctx, model.AdminUserListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	list := make([]v1.AdminUserItem, 0, len(out.List))
	for _, u := range out.List {
		list = append(list, v1.AdminUserItem{
			Id:       u.Id,
			Username: u.Username,
			Nickname: u.Nickname,
			Email:    u.Email,
			Phone:    u.Phone,
			Status:   u.Status,
		})
	}
	return &v1.AdminUserListRes{
		List:  list,
		Total: out.Total,
		Page:  out.Page,
	}, nil
}

func (c *Controller) AdminUserStatus(ctx context.Context, req *v1.AdminUserStatusReq) (res *v1.AdminUserStatusRes, err error) {
	err = service.Admin().UserSetStatus(ctx, model.AdminUserStatusInput{
		Id:     req.Id,
		Status: req.Status,
	})
	return
}

func (c *Controller) AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteReq) (res *v1.AdminUserDeleteRes, err error) {
	err = service.Admin().UserDelete(ctx, req.Id)
	return
}
