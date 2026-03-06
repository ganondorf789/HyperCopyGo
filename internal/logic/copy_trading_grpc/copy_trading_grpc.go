package copy_trading_grpc

import (
	"context"
	"fmt"

	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterCopyTradingGrpc(&sCopyTradingGrpc{})
}

type sCopyTradingGrpc struct{}

func (s *sCopyTradingGrpc) GetAutoCopyTradingList(ctx context.Context, appId, appSecret string) (list []model.CopyTradingItem, err error) {
	var appKey entity.UserAppKey
	err = dao.UserAppKey.Ctx(ctx).
		Where("app_id = ? AND app_secret = ? AND status = ?", appId, appSecret, consts.UserStatusEnabled).
		Scan(&appKey)
	if err != nil {
		return nil, err
	}
	if appKey.Id == 0 {
		return nil, fmt.Errorf("invalid app credentials")
	}

	var items []entity.CopyTrading
	err = dao.CopyTrading.Ctx(ctx).
		Where("user_id = ?", appKey.UserId).
		Where("follow_type = ?", consts.FollowTypeAuto).
		OrderDesc(dao.CopyTrading.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list = make([]model.CopyTradingItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToItem(e))
	}
	return list, nil
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
