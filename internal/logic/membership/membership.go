package membership

import (
	"context"
	"fmt"

	v1 "demo/api/membership/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterMembership(&sMembership{})
}

type sMembership struct{}

func (s *sMembership) Create(ctx context.Context, in v1.MembershipCreateReq) (res *v1.MembershipCreateRes, err error) {
	id, err := dao.Membership.Ctx(ctx).Data(do.Membership{
		UserId:   in.UserId,
		Level:    in.Level,
		StartAt:  in.StartAt,
		ExpireAt: in.ExpireAt,
		Status:   in.Status,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.MembershipCreateRes{Id: id}, nil
}

func (s *sMembership) Update(ctx context.Context, in v1.MembershipUpdateReq) error {
	count, err := dao.Membership.Ctx(ctx).Where(do.Membership{Id: in.Id}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("会员记录不存在")
	}

	_, err = dao.Membership.Ctx(ctx).
		Where(do.Membership{Id: in.Id}).
		Data(do.Membership{
			Level:    in.Level,
			StartAt:  in.StartAt,
			ExpireAt: in.ExpireAt,
			Status:   in.Status,
		}).
		Update()
	return err
}

func (s *sMembership) Delete(ctx context.Context, id int64) error {
	_, err := dao.Membership.Ctx(ctx).Where(do.Membership{Id: id}).Delete()
	return err
}

func (s *sMembership) Detail(ctx context.Context, id int64) (res *v1.MembershipDetailRes, err error) {
	var item entity.Membership
	err = dao.Membership.Ctx(ctx).Where(do.Membership{Id: id}).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("会员记录不存在")
	}
	return &v1.MembershipDetailRes{
		MembershipItem: entityToItem(item),
	}, nil
}

func (s *sMembership) List(ctx context.Context, in v1.MembershipListReq) (res *v1.MembershipListRes, err error) {
	m := dao.Membership.Ctx(ctx)
	if in.UserId > 0 {
		m = m.Where(do.Membership{UserId: in.UserId})
	}
	if in.Level >= 0 {
		m = m.Where(do.Membership{Level: in.Level})
	}
	if in.Status >= 0 {
		m = m.Where(do.Membership{Status: in.Status})
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.Membership
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.Membership.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.MembershipItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToItem(e))
	}

	return &v1.MembershipListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.Membership) model.MembershipItem {
	return model.MembershipItem{
		Id:        e.Id,
		UserId:    e.UserId,
		Level:     e.Level,
		StartAt:   e.StartAt,
		ExpireAt:  e.ExpireAt,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
