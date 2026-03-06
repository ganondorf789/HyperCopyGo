package copy_trade_config_grpc

import (
	"context"

	v1 "demo/api/copy_trade_config_grpc/v1"
	"demo/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedCopyTradeConfigServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCopyTradeConfigServiceServer(s.Server, &Controller{})
}

func (*Controller) GetAutoCopyTradeConfigList(ctx context.Context, req *v1.GetAutoCopyTradeConfigListReq) (res *v1.GetAutoCopyTradeConfigListRes, err error) {
	list, err := service.CopyTradeConfigGrpc().GetAutoCopyTradeConfigList(ctx, req.AppId, req.AppSecret)
	if err != nil {
		return nil, err
	}
	return &v1.GetAutoCopyTradeConfigListRes{
		List:  list,
		Total: int32(len(list)),
	}, nil
}

