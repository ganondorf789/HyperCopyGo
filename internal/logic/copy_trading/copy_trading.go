package copy_trading

import (
	"context"
	"fmt"

	v1 "demo/api/copy_trading/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterCopyTrading(&sCopyTrading{})
}

type sCopyTrading struct{}

func (s *sCopyTrading) Create(ctx context.Context, userId int64, in v1.CopyTradingCreateReq) (res *v1.CopyTradingCreateRes, err error) {
	switch in.FollowType {
	case consts.FollowTypeAuto: // 自动跟单：同一用户下 TargetWallet 不可重复
		count, err := dao.CopyTrading.Ctx(ctx).
			Where("user_id = ? AND target_wallet = ?", userId, in.TargetWallet).
			Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, fmt.Errorf("该目标钱包已存在跟单配置")
		}
	case consts.FollowTypeCondition: // 条件跟单：FollowSymbol 必填
		if in.FollowSymbol == "" {
			return nil, fmt.Errorf("条件跟单必须指定跟单币种")
		}
	}

	data := gconv.Map(in)
	data["user_id"] = userId

	id, err := dao.CopyTrading.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &v1.CopyTradingCreateRes{Id: id}, nil
}

func (s *sCopyTrading) Update(ctx context.Context, userId int64, in v1.CopyTradingUpdateReq) error {
	count, err := dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("跟单配置不存在")
	}

	data := gconv.Map(in)
	delete(data, "id")
	delete(data, "targetWallet")

	_, err = dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Data(data).
		OmitEmpty().
		Update()
	return err
}

func (s *sCopyTrading) Delete(ctx context.Context, userId int64, id int64) error {
	_, err := dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Delete()
	return err
}

func (s *sCopyTrading) Detail(ctx context.Context, userId int64, id int64) (res *v1.CopyTradingDetailRes, err error) {
	var item entity.CopyTrading
	err = dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
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
	m := dao.CopyTrading.Ctx(ctx).Where("user_id = ?", userId)
	if in.Status >= 0 {
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

	list := make([]model.CopyTradingItem, 0, len(items))
	for _, item := range items {
		list = append(list, entityToItem(item))
	}

	return &v1.CopyTradingListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func entityToItem(e entity.CopyTrading) model.CopyTradingItem {
	return model.CopyTradingItem{
		Id: e.Id,
		BaseCopyTrading: model.BaseCopyTrading{
			TargetWallet:                   e.TargetWallet,
			TargetWalletPlatform:           e.TargetWalletPlatform,
			Remark:                         e.Remark,
			FollowType:                     e.FollowType,
			FollowOnce:                     e.FollowOnce,
			PositionConditions:             e.PositionConditions,
			TraderConditions:               e.TraderConditions,
			TagAccountValue:                e.TagAccountValue,
			TagProfitScale:                 e.TagProfitScale,
			TagDirection:                   e.TagDirection,
			TagTradingRhythm:               e.TagTradingRhythm,
			TagProfitStatus:                e.TagProfitStatus,
			TagTradingStyles:               e.TagTradingStyles,
			TraderMetricPeriod:             e.TraderMetricPeriod,
			FollowMarginMode:               e.FollowMarginMode,
			FollowSymbol:                   e.FollowSymbol,
			Leverage:                       e.Leverage,
			MarginMode:                     e.MarginMode,
			FollowModel:                    e.FollowModel,
			FollowModelValue:               e.FollowModelValue,
			MinValue:                       e.MinValue,
			MaxValue:                       e.MaxValue,
			MaxMarginUsage:                 e.MaxMarginUsage,
			TpValue:                        e.TpValue,
			SlValue:                        e.SlValue,
			OptReverseFollowOrder:          e.OptReverseFollowOrder,
			OptFollowupDecrease:            e.OptFollowupDecrease,
			OptFollowupIncrease:            e.OptFollowupIncrease,
			OptForcedLiquidationProtection: e.OptForcedLiquidationProtection,
			OptPositionIncreaseOpening:     e.OptPositionIncreaseOpening,
			OptSlippageProtection:          e.OptSlippageProtection,
			SymbolListType:                 e.SymbolListType,
			SymbolList:                     e.SymbolList,
			MainWallet:                     e.MainWallet,
			MainWalletPlatform:             e.MainWalletPlatform,
		},
		Status:    int(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
