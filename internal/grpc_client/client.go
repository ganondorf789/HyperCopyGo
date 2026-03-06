package grpc_client

import (
	"context"

	pb "demo/api/copy_trading_grpc/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
)

type CopyTradingClient struct {
	conn      *grpc.ClientConn
	client    pb.CopyTradingServiceClient
	appId     string
	appSecret string
}

func NewCopyTradingClient(serviceName, appId, appSecret string) *CopyTradingClient {
	conn := grpcx.Client.MustNewGrpcClientConn(serviceName)
	return &CopyTradingClient{
		conn:      conn,
		client:    pb.NewCopyTradingServiceClient(conn),
		appId:     appId,
		appSecret: appSecret,
	}
}

func (c *CopyTradingClient) GetAutoCopyTradingList(ctx context.Context, page, pageSize int32) (*pb.GetAutoCopyTradingListRes, error) {
	res, err := c.client.GetAutoCopyTradingList(ctx, &pb.GetAutoCopyTradingListReq{
		AppId:     c.appId,
		AppSecret: c.appSecret,
		Page:      page,
		PageSize:  pageSize,
	})
	if err != nil {
		g.Log().Errorf(ctx, "GetAutoCopyTradingList rpc error: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *CopyTradingClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
