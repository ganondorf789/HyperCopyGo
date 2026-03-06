package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/copy_trade_config_grpc"
	"demo/internal/controller/wallet_grpc"
	"demo/internal/initialization"

	_ "demo/internal/logic/copy_trade_config_grpc"
	_ "demo/internal/logic/wallet_grpc"
)

var (
	Grpc = gcmd.Command{
		Name:  "grpc",
		Usage: "grpc",
		Brief: "start grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if err = initialization.InitConfig(); err != nil {
				return err
			}

			s := grpcx.Server.New()
			copy_trade_config_grpc.Register(s)
			wallet_grpc.Register(s)
			s.Run()
			return nil
		},
	}
)

