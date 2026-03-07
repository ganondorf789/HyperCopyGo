package server_management

import (
	"context"
	"fmt"

	v1 "demo/api/server_management/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterServerManagement(&sServerManagement{})
}

type sServerManagement struct{}

func (s *sServerManagement) Create(ctx context.Context, in v1.ServerManagementCreateReq) (res *v1.ServerManagementCreateRes, err error) {
	count, err := dao.ServerManagement.Ctx(ctx).
		Where("ip = ?", in.Ip).
		Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("该IP已存在")
	}

	id, err := dao.ServerManagement.Ctx(ctx).Data(g.Map{
		"ip":       in.Ip,
		"username": in.Username,
		"password": in.Password,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.ServerManagementCreateRes{Id: id}, nil
}

func (s *sServerManagement) Update(ctx context.Context, in v1.ServerManagementUpdateReq) error {
	count, err := dao.ServerManagement.Ctx(ctx).
		Where("id = ?", in.Id).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("服务器不存在")
	}

	data := g.Map{}
	if in.Ip != "" {
		data["ip"] = in.Ip
	}
	if in.Username != "" {
		data["username"] = in.Username
	}
	if in.Password != "" {
		data["password"] = in.Password
	}
	if len(data) == 0 {
		return nil
	}

	_, err = dao.ServerManagement.Ctx(ctx).
		Where("id = ?", in.Id).
		Data(data).
		Update()
	return err
}

func (s *sServerManagement) Delete(ctx context.Context, id int64) error {
	_, err := dao.ServerManagement.Ctx(ctx).
		Where("id = ?", id).
		Delete()
	return err
}

func (s *sServerManagement) Detail(ctx context.Context, id int64) (res *v1.ServerManagementDetailRes, err error) {
	var item entity.ServerManagement
	err = dao.ServerManagement.Ctx(ctx).
		Where("id = ?", id).
		Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("服务器不存在")
	}
	return &v1.ServerManagementDetailRes{
		ServerManagementItem: entityToItem(item),
	}, nil
}

func (s *sServerManagement) List(ctx context.Context, in v1.ServerManagementListReq) (res *v1.ServerManagementListRes, err error) {
	m := dao.ServerManagement.Ctx(ctx)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.ServerManagement
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.ServerManagement.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]v1.ServerManagementItem, 0, len(items))
	for _, item := range items {
		list = append(list, entityToItem(item))
	}

	return &v1.ServerManagementListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.ServerManagement) v1.ServerManagementItem {
	return v1.ServerManagementItem{
		Id:        e.Id,
		Ip:        e.Ip,
		Username:  e.Username,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
