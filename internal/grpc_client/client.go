package grpc_client

import (
	"context"

	pb "demo/api/copy_trade_config_grpc/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
)

type CopyTradeConfigClient struct {
	conn      *grpc.ClientConn
	client    pb.CopyTradeConfigServiceClient
	appId     string
	appSecret string
}

func NewCopyTradeConfigClient(serviceName, appId, appSecret string) *CopyTradeConfigClient {
	conn := grpcx.Client.MustNewGrpcClientConn(serviceName)
	return &CopyTradeConfigClient{
		conn:      conn,
		client:    pb.NewCopyTradeConfigServiceClient(conn),
		appId:     appId,
		appSecret: appSecret,
	}
}

func (c *CopyTradeConfigClient) GetAutoCopyTradeConfigList(ctx context.Context) (*pb.GetAutoCopyTradeConfigListRes, error) {
	res, err := c.client.GetAutoCopyTradeConfigList(ctx, &pb.GetAutoCopyTradeConfigListReq{
		AppId:     c.appId,
		AppSecret: c.appSecret,
	})
	if err != nil {
		g.Log().Errorf(ctx, "GetAutoCopyTradeConfigList rpc error: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *CopyTradeConfigClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

