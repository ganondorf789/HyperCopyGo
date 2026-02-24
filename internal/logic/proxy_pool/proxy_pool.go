package proxy_pool

import (
	"context"
	"fmt"

	v1 "demo/api/proxy_pool/v1"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterProxyPool(&sProxyPool{})
}

type sProxyPool struct{}

func (s *sProxyPool) Create(ctx context.Context, in v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error) {
	id, err := dao.ProxyPool.Ctx(ctx).Data(do.ProxyPool{
		Host:     in.Host,
		Port:     in.Port,
		Username: in.Username,
		Password: in.Password,
		Status:   1,
		Remark:   in.Remark,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.ProxyPoolCreateRes{Id: id}, nil
}

func (s *sProxyPool) Update(ctx context.Context, in v1.ProxyPoolUpdateReq) error {
	count, err := dao.ProxyPool.Ctx(ctx).Where(do.ProxyPool{Id: in.Id}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("代理不存在")
	}

	data := do.ProxyPool{}
	if in.Host != "" {
		data.Host = in.Host
	}
	if in.Port != 0 {
		data.Port = in.Port
	}
	if in.Username != "" {
		data.Username = in.Username
	}
	if in.Password != "" {
		data.Password = in.Password
	}
	if in.Status != nil {
		data.Status = *in.Status
	}
	if in.Remark != "" {
		data.Remark = in.Remark
	}

	_, err = dao.ProxyPool.Ctx(ctx).Where(do.ProxyPool{Id: in.Id}).Data(data).Update()
	return err
}

func (s *sProxyPool) Delete(ctx context.Context, id int64) error {
	_, err := dao.ProxyPool.Ctx(ctx).Where(do.ProxyPool{Id: id}).Delete()
	return err
}

func (s *sProxyPool) List(ctx context.Context, in v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error) {
	m := dao.ProxyPool.Ctx(ctx)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.ProxyPool
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.ProxyPool.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	return &v1.ProxyPoolListRes{
		List:  items,
		Total: total,
		Page:  in.Page,
	}, nil
}
