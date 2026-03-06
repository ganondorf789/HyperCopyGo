package wallet_grpc

import (
	"context"

	v1 "demo/api/wallet_grpc/v1"
	"demo/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedWalletServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterWalletServiceServer(s.Server, &Controller{})
}

func (*Controller) GetWalletList(ctx context.Context, req *v1.GetWalletListReq) (res *v1.GetWalletListRes, err error) {
	list, err := service.WalletGrpc().GetWalletList(ctx, req.AppId, req.AppSecret)
	if err != nil {
		return nil, err
	}
	return &v1.GetWalletListRes{List: list}, nil
}
