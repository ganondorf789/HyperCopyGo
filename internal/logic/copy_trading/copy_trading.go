package copy_trading

import (
	"context"
	"fmt"

	v1 "demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterCopyTrading(&sCopyTrading{})
}

type sCopyTrading struct{}

func (s *sCopyTrading) List(ctx context.Context, userId int64, in v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error) {
	m := dao.CopyTrading.Ctx(ctx).Where("user_id = ?", userId)
	if in.Status != "" {
		m = m.Where("status = ?", in.Status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.CopyTrading
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.CopyTrading.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	return &v1.CopyTradingListRes{
		List:  items,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sCopyTrading) Stop(ctx context.Context, userId int64, id int64) error {
	var item entity.CopyTrading
	err := dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Scan(&item)
	if err != nil {
		return err
	}
	if item.Id == 0 {
		return fmt.Errorf("跟单交易不存在")
	}
	if item.Status == consts.CopyTradingStatusStopped {
		return fmt.Errorf("该跟单交易已停止")
	}
	if item.Status != consts.CopyTradingStatusFollowing && item.Status != consts.CopyTradingStatusNotStarted {
		return fmt.Errorf("当前状态不允许停止")
	}

	_, err = dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Data("status", consts.CopyTradingStatusStopped).
		Update()
	return err
}
