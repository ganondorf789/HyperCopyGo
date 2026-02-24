package user

import (
	"context"

	v1 "demo/api/user/v1"
	"demo/internal/consts"
	"demo/internal/model"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	err = service.User().Register(ctx, model.UserRegisterInput{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

func (c *Controller) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	out, err := service.User().Login(ctx, model.UserLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserLoginRes{
		Token:  out.Token,
		Expire: out.Expire,
	}, nil
}

func (c *Controller) UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxUserIdKey).Int64()
	out, err := service.User().Profile(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &v1.UserProfileRes{
		Id:       out.Id,
		Username: out.Username,
		Nickname: out.Nickname,
		Avatar:   out.Avatar,
		Email:    out.Email,
		Phone:    out.Phone,
	}, nil
}
