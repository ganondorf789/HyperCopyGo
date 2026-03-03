package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/admin"
	appVersionCtrl "demo/internal/controller/app_version"
	copyTradingCtrl "demo/internal/controller/copy_trading"
	myTrackWalletCtrl "demo/internal/controller/my_track_wallet"
	positionCtrl "demo/internal/controller/position"
	proxyPoolCtrl "demo/internal/controller/proxy_pool"
	"demo/internal/controller/user"
	walletCtrl "demo/internal/controller/wallet"
	"demo/internal/initialization"
	"demo/internal/middleware"
	proxyPool "demo/internal/proxy_pool"

	_ "demo/internal/logic/admin"
	_ "demo/internal/logic/app_version"
	_ "demo/internal/logic/copy_trading"
	_ "demo/internal/logic/my_track_wallet"
	_ "demo/internal/logic/position"
	_ "demo/internal/logic/proxy_pool"
	_ "demo/internal/logic/user"
	_ "demo/internal/logic/wallet"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化配置
			if err = initialization.InitConfig(); err != nil {
				return err
			}

			// 初始化代理池（从数据库加载）
			if err = proxyPool.Reload(); err != nil {
				return err
			}

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(
						middleware.Auth,
						middleware.AdminAuth,
					)
					group.Bind(
						user.New(),
						admin.New(),
						appVersionCtrl.NewV1(),
						copyTradingCtrl.New(),
						walletCtrl.New(),
						myTrackWalletCtrl.New(),
						positionCtrl.New(),
						proxyPoolCtrl.New(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
