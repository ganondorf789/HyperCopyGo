package proxy_pool

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	v1 "demo/api/proxy_pool/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterProxyPool(&sProxyPool{})
}

type sProxyPool struct{}

func (s *sProxyPool) Create(ctx context.Context, in v1.ProxyPoolCreateReq) (res *v1.ProxyPoolCreateRes, err error) {
	id, err := dao.ProxyPools.Ctx(ctx).Data(entity.ProxyPools{
		Host:     in.Host,
		Port:     strconv.Itoa(in.Port),
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
	count, err := dao.ProxyPools.Ctx(ctx).Where("id = ?", in.Id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("代理不存在")
	}

	data := entity.ProxyPools{}
	if in.Host != "" {
		data.Host = in.Host
	}
	if in.Port != 0 {
		data.Port = strconv.Itoa(in.Port)
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

	_, err = dao.ProxyPools.Ctx(ctx).Where("id = ?", in.Id).Data(data).OmitEmpty().Update()
	return err
}

func (s *sProxyPool) Delete(ctx context.Context, id int64) error {
	_, err := dao.ProxyPools.Ctx(ctx).Where("id = ?", id).Delete()
	return err
}

func (s *sProxyPool) List(ctx context.Context, in v1.ProxyPoolListReq) (res *v1.ProxyPoolListRes, err error) {
	m := dao.ProxyPools.Ctx(ctx)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.ProxyPools
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.ProxyPools.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.ProxyPoolItem, 0, len(items))
	for _, e := range items {
		port, _ := strconv.Atoi(e.Port)
		list = append(list, model.ProxyPoolItem{
			Id: e.Id,
			BaseProxyPool: model.BaseProxyPool{
				Host:     e.Host,
				Port:     port,
				Username: e.Username,
				Password: e.Password,
				Remark:   e.Remark,
			},
			Status:    e.Status,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		})
	}

	return &v1.ProxyPoolListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sProxyPool) ImportFromCSV(ctx context.Context, in v1.ProxyPoolImportCSVReq) (res *v1.ProxyPoolImportCSVRes, err error) {
	f, err := in.File.Open()
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// skip header row: Host,Port,User,Pass
	if _, err = reader.Read(); err != nil {
		return nil, fmt.Errorf("读取CSV头部失败: %w", err)
	}

	var created, skipped int
	for {
		record, readErr := reader.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil || len(record) < 4 {
			skipped++
			continue
		}
		host := strings.TrimSpace(record[0])
		portStr := strings.TrimSpace(record[1])
		username := strings.TrimSpace(record[2])
		password := strings.TrimSpace(record[3])

		if host == "" || portStr == "" {
			skipped++
			continue
		}

		_, insertErr := dao.ProxyPools.Ctx(ctx).Data(entity.ProxyPools{
			Host:     host,
			Port:     portStr,
			Username: username,
			Password: password,
			Status:   1,
		}).InsertAndGetId()
		if insertErr != nil {
			skipped++
		} else {
			created++
		}
	}

	return &v1.ProxyPoolImportCSVRes{Created: created, Skipped: skipped}, nil
}
