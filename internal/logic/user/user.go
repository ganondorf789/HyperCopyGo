package user

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	v1 "demo/api/user/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/service"
	"demo/utility"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterUser(&sUser{})
}

type sUser struct{}

const verifyCodePrefix = "verify_code:"
const verifyCodeExpire = 5 * time.Minute

func (s *sUser) SendVerifyCode(ctx context.Context, in v1.SendVerifyCodeReq) error {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	redis := g.Redis()
	_, err := redis.Set(ctx, verifyCodePrefix+in.Email, code)
	if err != nil {
		return err
	}
	_, err = redis.Expire(ctx, verifyCodePrefix+in.Email, int64(verifyCodeExpire.Seconds()))
	if err != nil {
		return err
	}

	return service.Email().SendVerifyCode(ctx, in.Email, code)
}

func (s *sUser) Login(ctx context.Context, in v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	redis := g.Redis()
	cachedCode, err := redis.Get(ctx, verifyCodePrefix+in.Email)
	if err != nil {
		return nil, err
	}
	if cachedCode.IsEmpty() || cachedCode.String() != in.Code {
		return nil, fmt.Errorf("验证码错误或已过期")
	}

	_, _ = redis.Del(ctx, verifyCodePrefix+in.Email)

	var user entity.User
	err = dao.User.Ctx(ctx).
		Where("email = ?", in.Email).
		Scan(&user)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		id, err := dao.User.Ctx(ctx).Data(entity.User{
			Email:    in.Email,
			Username: in.Email,
			Status:   consts.UserStatusEnabled,
		}).InsertAndGetId()
		if err != nil {
			return nil, err
		}
		user.Id = id
	}

	if user.Status == consts.UserStatusDisabled {
		return nil, fmt.Errorf("账号已被禁用")
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
		Where("id = ?", userId).
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
