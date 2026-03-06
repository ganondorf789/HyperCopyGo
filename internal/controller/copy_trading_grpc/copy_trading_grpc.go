package copy_trading_grpc

import (
	"context"
	"fmt"

	v1 "demo/api/copy_trading_grpc/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/entity"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedCopyTradingServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCopyTradingServiceServer(s.Server, &Controller{})
}

func (*Controller) GetAutoCopyTradingList(ctx context.Context, req *v1.GetAutoCopyTradingListReq) (res *v1.GetAutoCopyTradingListRes, err error) {
	var appKey entity.UserAppKey
	err = dao.UserAppKey.Ctx(ctx).
		Where("app_id = ? AND app_secret = ? AND status = ?", req.AppId, req.AppSecret, consts.UserStatusEnabled).
		Scan(&appKey)
	if err != nil {
		return nil, err
	}
	if appKey.Id == 0 {
		return nil, fmt.Errorf("invalid app credentials")
	}

	page := int(req.Page)
	if page <= 0 {
		page = 1
	}
	pageSize := int(req.PageSize)
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	m := dao.CopyTrading.Ctx(ctx).
		Where("user_id = ?", appKey.UserId).
		Where("follow_type = ?", consts.FollowTypeAuto)

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.CopyTrading
	err = m.Page(page, pageSize).
		OrderDesc(dao.CopyTrading.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.CopyTradingItem, 0, len(items))
	for _, e := range items {
		list = append(list, entityToProto(e))
	}

	return &v1.GetAutoCopyTradingListRes{
		List:  list,
		Total: int32(total),
		Page:  int32(page),
	}, nil
}

func entityToProto(e entity.CopyTrading) *v1.CopyTradingItem {
	item := &v1.CopyTradingItem{
		Id:                             e.Id,
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
		Status:                         e.Status,
	}
	if e.CreatedAt != nil {
		item.CreatedAt = e.CreatedAt.String()
	}
	if e.UpdatedAt != nil {
		item.UpdatedAt = e.UpdatedAt.String()
	}
	return item
}
