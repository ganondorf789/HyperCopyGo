package copy_trading

import (
	"context"
	"fmt"

	v1 "demo/api/copy_trading/v1"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterCopyTrading(&sCopyTrading{})
}

type sCopyTrading struct{}

func (s *sCopyTrading) Create(ctx context.Context, userId int64, in v1.CopyTradingCreateReq) (res *v1.CopyTradingCreateRes, err error) {
	id, err := dao.CopyTrading.Ctx(ctx).Data(do.CopyTrading{
		UserId:                         userId,
		TargetWallet:                   in.TargetWallet,
		TargetWalletPlatform:           in.TargetWalletPlatform,
		Remark:                         in.Remark,
		Leverage:                       in.Leverage,
		MarginMode:                     in.MarginMode,
		FollowModel:                    in.FollowModel,
		FollowModelValue:               in.FollowModelValue,
		MinValue:                       in.MinValue,
		MaxValue:                       in.MaxValue,
		MaxMarginUsage:                 in.MaxMarginUsage,
		TpValue:                        in.TpValue,
		SlValue:                        in.SlValue,
		OptReverseFollowOrder:          in.OptReverseFollowOrder,
		OptFollowupDecrease:            in.OptFollowupDecrease,
		OptFollowupIncrease:            in.OptFollowupIncrease,
		OptForcedLiquidationProtection: in.OptForcedLiquidationProtection,
		OptPositionIncreaseOpening:     in.OptPositionIncreaseOpening,
		OptSlippageProtection:          in.OptSlippageProtection,
		SymbolListType:                 in.SymbolListType,
		SymbolList:                     in.SymbolList,
		MainWallet:                     in.MainWallet,
		MainWalletPlatform:             in.MainWalletPlatform,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.CopyTradingCreateRes{Id: id}, nil
}

func (s *sCopyTrading) Update(ctx context.Context, userId int64, in v1.CopyTradingUpdateReq) error {
	count, err := dao.CopyTrading.Ctx(ctx).
		Where(do.CopyTrading{Id: in.Id, UserId: userId}).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("跟单配置不存在")
	}

	_, err = dao.CopyTrading.Ctx(ctx).
		Where(do.CopyTrading{Id: in.Id, UserId: userId}).
		Data(do.CopyTrading{
			TargetWallet:                   in.TargetWallet,
			TargetWalletPlatform:           in.TargetWalletPlatform,
			Remark:                         in.Remark,
			Leverage:                       in.Leverage,
			MarginMode:                     in.MarginMode,
			FollowModel:                    in.FollowModel,
			FollowModelValue:               in.FollowModelValue,
			MinValue:                       in.MinValue,
			MaxValue:                       in.MaxValue,
			MaxMarginUsage:                 in.MaxMarginUsage,
			TpValue:                        in.TpValue,
			SlValue:                        in.SlValue,
			OptReverseFollowOrder:          in.OptReverseFollowOrder,
			OptFollowupDecrease:            in.OptFollowupDecrease,
			OptFollowupIncrease:            in.OptFollowupIncrease,
			OptForcedLiquidationProtection: in.OptForcedLiquidationProtection,
			OptPositionIncreaseOpening:     in.OptPositionIncreaseOpening,
			OptSlippageProtection:          in.OptSlippageProtection,
			SymbolListType:                 in.SymbolListType,
			SymbolList:                     in.SymbolList,
			MainWallet:                     in.MainWallet,
			MainWalletPlatform:             in.MainWalletPlatform,
			Status:                         in.Status,
		}).
		Update()
	return err
}

func (s *sCopyTrading) Delete(ctx context.Context, userId int64, id int64) error {
	_, err := dao.CopyTrading.Ctx(ctx).
		Where(do.CopyTrading{Id: id, UserId: userId}).
		Delete()
	return err
}

func (s *sCopyTrading) Detail(ctx context.Context, userId int64, id int64) (res *v1.CopyTradingDetailRes, err error) {
	var item entity.CopyTrading
	err = dao.CopyTrading.Ctx(ctx).
		Where(do.CopyTrading{Id: id, UserId: userId}).
		Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("跟单配置不存在")
	}
	return &v1.CopyTradingDetailRes{
		CopyTradingItem: entityToItem(item),
	}, nil
}

func (s *sCopyTrading) List(ctx context.Context, userId int64, in v1.CopyTradingListReq) (res *v1.CopyTradingListRes, err error) {
	m := dao.CopyTrading.Ctx(ctx).Where(do.CopyTrading{UserId: userId})
	if in.Status >= 0 {
		m = m.Where(do.CopyTrading{Status: in.Status})
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

	list := make([]v1.CopyTradingItem, 0, len(items))
	for _, item := range items {
		list = append(list, entityToItem(item))
	}

	return &v1.CopyTradingListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.CopyTrading) v1.CopyTradingItem {
	return v1.CopyTradingItem{
		Id:                 e.Id,
		BaseCopyTrading:    e.BaseCopyTrading,
		Status:             e.Status,
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
	}
}
