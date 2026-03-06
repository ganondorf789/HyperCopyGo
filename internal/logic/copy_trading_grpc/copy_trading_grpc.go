package copy_trading_grpc

import (
	"context"
	"fmt"

	v1 "demo/api/copy_trading_grpc/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterCopyTradingGrpc(&sCopyTradingGrpc{})
}

type sCopyTradingGrpc struct{}

func (s *sCopyTradingGrpc) GetCopyTradingDetail(ctx context.Context, appId, appSecret string, id int64) (*v1.CopyTradingItem, error) {
	userId, err := validateAppKey(ctx, appId, appSecret)
	if err != nil {
		return nil, err
	}

	var item entity.CopyTrading
	err = dao.CopyTrading.Ctx(ctx).
		Where("id = ? AND user_id = ?", id, userId).
		Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, fmt.Errorf("copy trading record not found")
	}

	return entityToProto(item), nil
}

func (s *sCopyTradingGrpc) GetCopyTradingList(ctx context.Context, appId, appSecret string, copyTradingId int64) (list []*v1.CopyTradingItem, err error) {
	userId, err := validateAppKey(ctx, appId, appSecret)
	if err != nil {
		return nil, err
	}

	m := dao.CopyTrading.Ctx(ctx).Where("user_id = ?", userId)
	if copyTradingId > 0 {
		m = m.Where("copy_trading_id = ?", copyTradingId)
	}

	var items []entity.CopyTrading
	err = m.OrderDesc(dao.CopyTrading.Columns().Id).Scan(&items)
	if err != nil {
		return nil, err
	}

	list = make([]*v1.CopyTradingItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToProto(e))
	}
	return list, nil
}

func validateAppKey(ctx context.Context, appId, appSecret string) (int64, error) {
	var appKey entity.UserAppKey
	err := dao.UserAppKey.Ctx(ctx).
		Where("app_id = ? AND app_secret = ? AND status = ?", appId, appSecret, consts.UserStatusEnabled).
		Scan(&appKey)
	if err != nil {
		return 0, err
	}
	if appKey.Id == 0 {
		return 0, fmt.Errorf("invalid app credentials")
	}
	return appKey.UserId, nil
}

func entityToProto(e entity.CopyTrading) *v1.CopyTradingItem {
	item := &v1.CopyTradingItem{
		Id:                             e.Id,
		CopyTradingId:                  e.CopyTradingId,
		UserId:                         e.UserId,
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
		CopyTradingStatus:              e.CopyTradingStatus,
		TraderAddress:                  e.TraderAddress,
		TraderCoin:                     e.TraderCoin,
		TraderSzi:                      e.TraderSzi,
		TraderLeverageType:             e.TraderLeverageType,
		TraderLeverage:                 e.TraderLeverage,
		TraderEntryPx:                  e.TraderEntryPx,
		TraderPositionValue:            e.TraderPositionValue,
		ExecuteStatus:                  e.ExecuteStatus,
		OrderStatus:                    e.OrderStatus,
		ErrorMsg:                       e.ErrorMsg,
	}
	if e.CopyTradingCreatedAt != nil {
		item.CopyTradingCreatedAt = e.CopyTradingCreatedAt.String()
	}
	if e.CopyTradingUpdatedAt != nil {
		item.CopyTradingUpdatedAt = e.CopyTradingUpdatedAt.String()
	}
	if e.CreatedAt != nil {
		item.CreatedAt = e.CreatedAt.String()
	}
	if e.UpdatedAt != nil {
		item.UpdatedAt = e.UpdatedAt.String()
	}
	return item
}
