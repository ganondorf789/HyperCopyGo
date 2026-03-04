package admin

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"

	v1 "demo/api/admin/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
	"demo/utility"
)

func init() {
	service.RegisterAdmin(&sAdmin{})
}

type sAdmin struct{}

func (s *sAdmin) Init(ctx context.Context, in v1.AdminInitReq) (res *v1.AdminInitRes, err error) {
	count, err := dao.Admin.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("系统已存在管理员，无法重复初始化")
	}

	id, err := dao.Admin.Ctx(ctx).Data(do.Admin{
		Username: in.Username,
		Password: encryptPassword(in.Password),
		Realname: in.Realname,
		Role:     consts.RoleSuperAdmin,
		Status:   consts.UserStatusEnabled,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.AdminInitRes{Id: id}, nil
}

func (s *sAdmin) Login(ctx context.Context, in v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	var admin entity.Admin
	err = dao.Admin.Ctx(ctx).
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

	token, expire, err := utility.GenerateToken(admin.Id, consts.UserTypeAdmin)
	if err != nil {
		return nil, err
	}
	return &v1.AdminLoginRes{Token: token, Expire: expire}, nil
}

func (s *sAdmin) Profile(ctx context.Context, adminId int64) (res *v1.AdminProfileRes, err error) {
	var admin entity.Admin
	err = dao.Admin.Ctx(ctx).
		Where(do.Admin{Id: adminId}).
		Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.Id == 0 {
		return nil, fmt.Errorf("管理员不存在")
	}
	return &v1.AdminProfileRes{
		Id:       admin.Id,
		Username: admin.Username,
		Realname: admin.Realname,
		Role:     admin.Role,
	}, nil
}

func (s *sAdmin) UserList(ctx context.Context, in v1.AdminUserListReq) (res *v1.AdminUserListRes, err error) {
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

	list := make([]v1.AdminUserItem, 0, len(users))
	for _, u := range users {
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
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sAdmin) UserSetStatus(ctx context.Context, in v1.AdminUserStatusReq) error {
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
