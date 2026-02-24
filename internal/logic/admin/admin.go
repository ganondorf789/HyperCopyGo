package admin

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
	service.RegisterAdmin(&sAdmin{})
}

type sAdmin struct{}

func (s *sAdmin) Login(ctx context.Context, in model.AdminLoginInput) (*model.TokenOutput, error) {
	var admin entity.Admin
	err := dao.Admin.Ctx(ctx).
		Where(do.Admin{Username: in.Username}).
		Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.Id == 0 {
		return nil, fmt.Errorf("管理员不存在")
	}
	if admin.Status == consts.UserStatusDisabled {
		return nil, fmt.Errorf("账号已被禁用")
	}
	if admin.Password != encryptPassword(in.Password) {
		return nil, fmt.Errorf("密码错误")
	}

	token, expire, err := middleware.GenerateToken(admin.Id, consts.UserTypeAdmin)
	if err != nil {
		return nil, err
	}
	return &model.TokenOutput{Token: token, Expire: expire}, nil
}

func (s *sAdmin) Profile(ctx context.Context, adminId int64) (*model.AdminInfoOutput, error) {
	var admin entity.Admin
	err := dao.Admin.Ctx(ctx).
		Where(do.Admin{Id: adminId}).
		Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.Id == 0 {
		return nil, fmt.Errorf("管理员不存在")
	}
	return &model.AdminInfoOutput{
		Id:       admin.Id,
		Username: admin.Username,
		Realname: admin.Realname,
		Role:     admin.Role,
		Status:   admin.Status,
	}, nil
}

func (s *sAdmin) UserList(ctx context.Context, in model.AdminUserListInput) (*model.AdminUserListOutput, error) {
	m := dao.User.Ctx(ctx)
	if in.Status >= 0 {
		m = m.Where(do.User{Status: in.Status})
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var users []entity.User
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.User.Columns().Id).
		Scan(&users)
	if err != nil {
		return nil, err
	}

	list := make([]model.UserInfoOutput, 0, len(users))
	for _, u := range users {
		list = append(list, model.UserInfoOutput{
			Id:       u.Id,
			Username: u.Username,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Email:    u.Email,
			Phone:    u.Phone,
			Status:   u.Status,
		})
	}

	return &model.AdminUserListOutput{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sAdmin) UserSetStatus(ctx context.Context, in model.AdminUserStatusInput) error {
	_, err := dao.User.Ctx(ctx).
		Where(do.User{Id: in.Id}).
		Data(do.User{Status: in.Status}).
		Update()
	return err
}

func (s *sAdmin) UserDelete(ctx context.Context, id int64) error {
	_, err := dao.User.Ctx(ctx).
		Where(do.User{Id: id}).
		Delete()
	return err
}

func encryptPassword(password string) string {
	return gmd5.MustEncryptString(password + "HyperCopyGo")
}
