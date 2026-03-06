package copy_trading_grpc

import (
	"context"

	v1 "demo/api/copy_trading_grpc/v1"
	"demo/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedCopyTradingServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCopyTradingServiceServer(s.Server, &Controller{})
}

func (*Controller) GetAutoCopyTradingList(ctx context.Context, req *v1.GetAutoCopyTradingListReq) (res *v1.GetAutoCopyTradingListRes, err error) {
	items, err := service.CopyTradingGrpc().GetAutoCopyTradingList(ctx, req.AppId, req.AppSecret)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.CopyTradingItem, 0, len(items))
	for _, e := range items {
		list = append(list, &v1.CopyTradingItem{
			Id:                             e.Id,
			UserId:                         0,
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
			Status:                         int64(e.Status),
		})
		if e.CreatedAt != nil {
			list[len(list)-1].CreatedAt = e.CreatedAt.String()
		}
		if e.UpdatedAt != nil {
			list[len(list)-1].UpdatedAt = e.UpdatedAt.String()
		}
	}

	return &v1.GetAutoCopyTradingListRes{
		List:  list,
		Total: int32(len(list)),
	}, nil
}
