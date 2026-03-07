package copy_trade_config

import (
	"context"
	"fmt"

	v1 "demo/api/copy_trade_config/v1"
	"demo/internal/consts"
	"demo/internal/cron_jobs"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterCopyTradeConfig(&sCopyTradeConfig{})
}

type sCopyTradeConfig struct{}

func (s *sCopyTradeConfig) Create(ctx context.Context, userId int64, in v1.CopyTradeConfigCreateReq) (res *v1.CopyTradeConfigCreateRes, err error) {
	switch in.FollowType {
	case consts.FollowTypeAuto: // 自动跟单：同一用户下 TargetWallet 不可重复
		count, err := dao.CopyTradeConfig.Ctx(ctx).
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

	id, err := dao.CopyTradeConfig.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	cron_jobs.TriggerAddressDispatch()
	return &v1.CopyTradeConfigCreateRes{Id: id}, nil
}

func (s *sCopyTradeConfig) Update(ctx context.Context, userId int64, in v1.CopyTradeConfigUpdateReq) error {
	count, err := dao.CopyTradeConfig.Ctx(ctx).
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

	_, err = dao.CopyTradeConfig.Ctx(ctx).
		Where("id = ? AND user_id = ?", in.Id, userId).
		Data(data).
		OmitEmpty().
		Update()
	if err != nil {
		return err
	}
	cron_jobs.TriggerAddressDispatch()
	return nil
}

func (s *sCopyTradeConfig) Delete(ctx context.Context, userId int64, id int64) error {
	_, err := dao.CopyTradeConfig.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Delete()
	if err != nil {
		return err
	}
	cron_jobs.TriggerAddressDispatch()
	return nil
}

func (s *sCopyTradeConfig) Detail(ctx context.Context, userId int64, id int64) (res *v1.CopyTradeConfigDetailRes, err error) {
	var item entity.CopyTradeConfig
	err = dao.CopyTradeConfig.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("跟单配置不存在")
	}
	return &v1.CopyTradeConfigDetailRes{
		CopyTradingItem: entityToItem(item),
	}, nil
}

func (s *sCopyTradeConfig) List(ctx context.Context, userId int64, in v1.CopyTradeConfigListReq) (res *v1.CopyTradeConfigListRes, err error) {
	m := dao.CopyTradeConfig.Ctx(ctx).Where("user_id = ?", userId)
	if in.Status >= 0 {
		m = m.Where("status = ?", in.Status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.CopyTradeConfig
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.CopyTradeConfig.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]model.CopyTradeConfig, 0, len(items))
	for _, item := range items {
		list = append(list, entityToItem(item))
	}

	return &v1.CopyTradeConfigListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sCopyTradeConfig) RecordList(ctx context.Context, userId int64, in v1.CopyTradeConfigRecordListReq) (res *v1.CopyTradeConfigRecordListRes, err error) {
	var records []entity.CopyTradeRecord
	err = dao.CopyTradeRecord.Ctx(ctx).
		Where("user_id = ? AND address = ?", userId, in.Address).
		OrderDesc(dao.CopyTradeRecord.Columns().Id).
		Scan(&records)
	if err != nil {
		return nil, err
	}

	list := make([]model.CopyTradeRecordItem, 0, len(records))
	for _, r := range records {
		list = append(list, model.CopyTradeRecordItem{
			Id:            r.Id,
			Address:       r.Address,
			Coin:          r.Coin,
			Direction:     r.Direction,
			Size:          r.Size,
			Price:         r.Price,
			ClosedPnl:     r.ClosedPnl,
			ExecuteStatus: r.ExecuteStatus,
			OrderStatus:   r.OrderStatus,
			ErrorMsg:      r.ErrorMsg,
			TradeTime:     r.TradeTime,
			CreatedAt:     r.CreatedAt,
		})
	}

	return &v1.CopyTradeConfigRecordListRes{List: list}, nil
}

func entityToItem(e entity.CopyTradeConfig) model.CopyTradeConfig {
	return model.CopyTradeConfig{
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




