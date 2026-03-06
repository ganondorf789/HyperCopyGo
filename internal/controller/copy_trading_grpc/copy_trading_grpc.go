package copy_trading_grpc

import (
	"context"

	v1 "demo/api/copy_trading_grpc/v1"
	"demo/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCopyTradingServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCopyTradingServiceServer(s.Server, &Controller{})
}

func (*Controller) GetCopyTradingDetail(ctx context.Context, req *v1.GetCopyTradingDetailReq) (res *v1.GetCopyTradingDetailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetCopyTradingList(ctx context.Context, req *v1.GetCopyTradingListReq) (res *v1.GetCopyTradingListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateCopyTrading(ctx context.Context, req *v1.CreateCopyTradingReq) (res *v1.CreateCopyTradingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SendCopyTradingNotification(ctx context.Context, req *v1.SendCopyTradingNotificationReq) (res *v1.SendCopyTradingNotificationRes, err error) {
	id, err := service.CopyTradingGrpc().SendCopyTradingNotification(ctx, req.AppId, req.AppSecret, req)
	if err != nil {
		return nil, err
	}
	return &v1.SendCopyTradingNotificationRes{Id: id}, nil
}
