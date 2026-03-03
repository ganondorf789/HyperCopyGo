package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"

	v1 "demo/api/user/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
	"demo/utility"
)

func init() {
	service.RegisterUser(&sUser{})
}

type sUser struct{}

func (s *sUser) Register(ctx context.Context, in v1.UserRegisterReq) error {
	count, err := dao.User.Ctx(ctx).
		Where(do.User{Username: in.Username}).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("用户名已存在")
	}

	_, err = dao.User.Ctx(ctx).Data(do.User{
		Username: in.Username,
		Password: encryptPassword(in.Password),
		Nickname: in.Nickname,
	}).Insert()
	return err
}

func (s *sUser) Login(ctx context.Context, in v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var user entity.User
	err = dao.User.Ctx(ctx).
		Where(do.User{Username: in.Username}).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	if user.Status == consts.UserStatusDisabled {
		return nil, fmt.Errorf("账号已被禁用")
	}
	if user.Password != encryptPassword(in.Password) {
		return nil, fmt.Errorf("密码错误")
	}

	token, expire, err := utility.GenerateToken(user.Id, consts.UserTypeUser)
	if err != nil {
		return nil, err
	}
	return &v1.UserLoginRes{Token: token, Expire: expire}, nil
}

func (s *sUser) Profile(ctx context.Context, userId int64) (res *v1.UserProfileRes, err error) {
	var user entity.User
	err = dao.User.Ctx(ctx).
		Where(do.User{Id: userId}).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	return &v1.UserProfileRes{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}

func encryptPassword(password string) string {
	return gmd5.MustEncryptString(password + "HyperCopyGo")
}
