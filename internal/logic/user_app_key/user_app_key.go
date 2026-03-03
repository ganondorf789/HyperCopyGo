package user_app_key

import (
	"context"
	"fmt"

	v1 "demo/api/user_app_key/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
)

func init() {
	service.RegisterUserAppKey(&sUserAppKey{})
}

type sUserAppKey struct{}

func (s *sUserAppKey) Create(ctx context.Context, in v1.UserAppKeyCreateReq) (res *v1.UserAppKeyCreateRes, err error) {
	appId := gstr.ToUpper(grand.S(16))
	appSecret := grand.S(32)

	id, err := dao.UserAppKey.Ctx(ctx).Data(do.UserAppKey{
		UserId:    in.UserId,
		AppId:     appId,
		AppSecret: appSecret,
		Remark:    in.Remark,
		ExpireAt:  in.ExpireAt,
		Status:    in.Status,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.UserAppKeyCreateRes{
		Id:        id,
		AppId:     appId,
		AppSecret: appSecret,
	}, nil
}

func (s *sUserAppKey) Update(ctx context.Context, in v1.UserAppKeyUpdateReq) error {
	count, err := dao.UserAppKey.Ctx(ctx).Where(do.UserAppKey{Id: in.Id}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("AppKey不存在")
	}

	_, err = dao.UserAppKey.Ctx(ctx).
		Where(do.UserAppKey{Id: in.Id}).
		Data(do.UserAppKey{
			Remark:   in.Remark,
			ExpireAt: in.ExpireAt,
			Status:   in.Status,
		}).
		Update()
	return err
}

func (s *sUserAppKey) Delete(ctx context.Context, id int64) error {
	_, err := dao.UserAppKey.Ctx(ctx).Where(do.UserAppKey{Id: id}).Delete()
	return err
}

func (s *sUserAppKey) Detail(ctx context.Context, id int64) (res *v1.UserAppKeyDetailRes, err error) {
	var item entity.UserAppKey
	err = dao.UserAppKey.Ctx(ctx).Where(do.UserAppKey{Id: id}).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("AppKey不存在")
	}
	return &v1.UserAppKeyDetailRes{
		UserAppKeyItem: entityToItem(item),
	}, nil
}

func (s *sUserAppKey) List(ctx context.Context, in v1.UserAppKeyListReq) (res *v1.UserAppKeyListRes, err error) {
	m := dao.UserAppKey.Ctx(ctx)
	if in.UserId > 0 {
		m = m.Where(do.UserAppKey{UserId: in.UserId})
	}
	if in.Status >= 0 {
		m = m.Where(do.UserAppKey{Status: in.Status})
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.UserAppKey
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.UserAppKey.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.UserAppKeyItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToItem(e))
	}

	return &v1.UserAppKeyListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.UserAppKey) model.UserAppKeyItem {
	return model.UserAppKeyItem{
		Id:        e.Id,
		UserId:    e.UserId,
		AppId:     e.AppId,
		AppSecret: e.AppSecret,
		Remark:    e.Remark,
		ExpireAt:  e.ExpireAt,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
