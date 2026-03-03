package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/admin"
	appVersionCtrl "demo/internal/controller/app_version"
	copyTradingCtrl "demo/internal/controller/copy_trading"
	cronTaskCtrl "demo/internal/controller/cron_task"
	membershipCtrl "demo/internal/controller/membership"
	myTrackWalletCtrl "demo/internal/controller/my_track_wallet"
	notificationCtrl "demo/internal/controller/notification"
	positionCtrl "demo/internal/controller/position"
	proxyPoolCtrl "demo/internal/controller/proxy_pool"
	"demo/internal/controller/user"
	userAppKeyCtrl "demo/internal/controller/user_app_key"
	walletCtrl "demo/internal/controller/wallet"
	"demo/internal/initialization"
	"demo/internal/middleware"
	proxyPool "demo/internal/proxy_pool"

	_ "demo/internal/logic/admin"
	_ "demo/internal/logic/app_version"
	_ "demo/internal/logic/copy_trading"
	_ "demo/internal/logic/cron_task"
	_ "demo/internal/logic/membership"
	_ "demo/internal/logic/my_track_wallet"
	_ "demo/internal/logic/notification"
	_ "demo/internal/logic/position"
	_ "demo/internal/logic/proxy_pool"
	_ "demo/internal/logic/user"
	_ "demo/internal/logic/user_app_key"
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
						cronTaskCtrl.NewV1(),
						copyTradingCtrl.New(),
						walletCtrl.New(),
						membershipCtrl.NewV1(),
						myTrackWalletCtrl.New(),
						notificationCtrl.NewV1(),
						positionCtrl.New(),
						proxyPoolCtrl.New(),
						userAppKeyCtrl.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
