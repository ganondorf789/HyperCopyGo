package app_version

import (
	"context"
	"fmt"

	v1 "demo/api/app_version/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterAppVersion(&sAppVersion{})
}

type sAppVersion struct{}

func (s *sAppVersion) Create(ctx context.Context, in v1.AppVersionCreateReq) (res *v1.AppVersionCreateRes, err error) {
	id, err := dao.AppVersion.Ctx(ctx).Data(do.AppVersion{
		Platform:       in.Platform,
		VersionName:    in.VersionName,
		VersionCode:    in.VersionCode,
		DownloadUrl:    in.DownloadUrl,
		ChangeLog:      in.ChangeLog,
		ForceUpdate:    in.ForceUpdate,
		MinVersionCode: in.MinVersionCode,
		Status:         in.Status,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.AppVersionCreateRes{Id: id}, nil
}

func (s *sAppVersion) Update(ctx context.Context, in v1.AppVersionUpdateReq) error {
	count, err := dao.AppVersion.Ctx(ctx).Where(do.AppVersion{Id: in.Id}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("版本记录不存在")
	}

	_, err = dao.AppVersion.Ctx(ctx).
		Where(do.AppVersion{Id: in.Id}).
		Data(do.AppVersion{
			Platform:       in.Platform,
			VersionName:    in.VersionName,
			VersionCode:    in.VersionCode,
			DownloadUrl:    in.DownloadUrl,
			ChangeLog:      in.ChangeLog,
			ForceUpdate:    in.ForceUpdate,
			MinVersionCode: in.MinVersionCode,
			Status:         in.Status,
		}).
		Update()
	return err
}

func (s *sAppVersion) Delete(ctx context.Context, id int64) error {
	_, err := dao.AppVersion.Ctx(ctx).Where(do.AppVersion{Id: id}).Delete()
	return err
}

func (s *sAppVersion) List(ctx context.Context, in v1.AppVersionListReq) (res *v1.AppVersionListRes, err error) {
	m := dao.AppVersion.Ctx(ctx)
	if in.Platform != "" {
		m = m.Where(do.AppVersion{Platform: in.Platform})
	}
	if in.Status >= 0 {
		m = m.Where(do.AppVersion{Status: in.Status})
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.AppVersion
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.AppVersion.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]v1.AppVersionItem, 0, len(items))
	for _, e := range items {
		list = append(list, v1.AppVersionItem{
			Id: e.Id,
			BaseAppVersion: model.BaseAppVersion{
				Platform:       e.Platform,
				VersionName:    e.VersionName,
				VersionCode:    e.VersionCode,
				DownloadUrl:    e.DownloadUrl,
				ChangeLog:      e.ChangeLog,
				ForceUpdate:    e.ForceUpdate,
				MinVersionCode: e.MinVersionCode,
			},
			Status:    e.Status,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		})
	}

	return &v1.AppVersionListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sAppVersion) Check(ctx context.Context, in v1.AppVersionCheckReq) (res *v1.AppVersionCheckRes, err error) {
	var latest entity.AppVersion
	err = dao.AppVersion.Ctx(ctx).
		Where(do.AppVersion{Platform: in.Platform, Status: 1}).
		OrderDesc(dao.AppVersion.Columns().VersionCode).
		Scan(&latest)
	if err != nil {
		return nil, err
	}

	if latest.Id == 0 || latest.VersionCode <= in.VersionCode {
		return &v1.AppVersionCheckRes{HasUpdate: false}, nil
	}

	forceUpdate := latest.ForceUpdate == 1 || in.VersionCode < latest.MinVersionCode
	return &v1.AppVersionCheckRes{
		HasUpdate:   true,
		ForceUpdate: forceUpdate,
		VersionName: latest.VersionName,
		VersionCode: latest.VersionCode,
		DownloadUrl: latest.DownloadUrl,
		ChangeLog:   latest.ChangeLog,
	}, nil
}
