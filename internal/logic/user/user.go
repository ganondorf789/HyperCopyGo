package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"

	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/middleware"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterUser(&sUser{})
}

type sUser struct{}

func (s *sUser) Register(ctx context.Context, in model.UserRegisterInput) error {
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

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (*model.TokenOutput, error) {
	var user entity.User
	err := dao.User.Ctx(ctx).
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

	token, expire, err := middleware.GenerateToken(user.Id, consts.UserTypeUser)
	if err != nil {
		return nil, err
	}
	return &model.TokenOutput{Token: token, Expire: expire}, nil
}

func (s *sUser) Profile(ctx context.Context, userId int64) (*model.UserInfoOutput, error) {
	var user entity.User
	err := dao.User.Ctx(ctx).
		Where(do.User{Id: userId}).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	return &model.UserInfoOutput{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Phone:    user.Phone,
		Status:   user.Status,
	}, nil
}

func encryptPassword(password string) string {
	return gmd5.MustEncryptString(password + "HyperCopyGo")
}
