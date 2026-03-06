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
	list, err := service.CopyTradingGrpc().GetAutoCopyTradingList(ctx, req.AppId, req.AppSecret)
	if err != nil {
		return nil, err
	}
	return &v1.GetAutoCopyTradingListRes{
		List:  list,
		Total: int32(len(list)),
	}, nil
}
