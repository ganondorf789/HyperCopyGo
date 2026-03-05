package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	adminCtrl "demo/internal/controller/admin"
	appVersionCtrl "demo/internal/controller/app_version"
	copyTradingCtrl "demo/internal/controller/copy_trading"
	cronTaskCtrl "demo/internal/controller/cron_task"
	leaderboardCtrl "demo/internal/controller/leaderboard"
	membershipCtrl "demo/internal/controller/membership"
	myTrackWalletCtrl "demo/internal/controller/my_track_wallet"
	notificationCtrl "demo/internal/controller/notification"
	positionCtrl "demo/internal/controller/position"
	proxyPoolCtrl "demo/internal/controller/proxy_pool"
	completedTradesCtrl "demo/internal/controller/completed_trades"
	traderCtrl "demo/internal/controller/trader"
	traderPerformanceCtrl "demo/internal/controller/trader_performance"
	traderPositionsCtrl "demo/internal/controller/trader_positions"
	userCtrl "demo/internal/controller/user"
	userAppKeyCtrl "demo/internal/controller/user_app_key"
	walletCtrl "demo/internal/controller/wallet"
	wsCtrl "demo/internal/controller/ws"
	"demo/internal/initialization"
	"demo/internal/middleware"
	proxyPool "demo/internal/proxy_pool"
	"demo/internal/service"
	"demo/internal/subscriber"

	_ "demo/internal/cron_jobs"
	_ "demo/internal/logic/completed_trades"
	_ "demo/internal/logic/trader_performance"
	_ "demo/internal/logic/admin"
	_ "demo/internal/logic/email"
	_ "demo/internal/logic/app_version"
	_ "demo/internal/logic/copy_trading"
	_ "demo/internal/logic/cron_task"
	_ "demo/internal/logic/leaderboard"
	_ "demo/internal/logic/membership"
	_ "demo/internal/logic/my_track_wallet"
	_ "demo/internal/logic/notification"
	_ "demo/internal/logic/position"
	_ "demo/internal/logic/proxy_pool"
	_ "demo/internal/logic/trader"
	_ "demo/internal/logic/trader_positions"
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

			// 从数据库加载并启动定时任务
			service.CronTask().StartAll(ctx)

			// 启动 Redis 订阅（new_positions / market_alert）
			subscriber.Start(ctx)

			s := g.Server()

			// WebSocket 路由（不使用 MiddlewareHandlerResponse，避免干扰 WebSocket 升级）
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Bind(wsCtrl.New())
			})

			// REST API 路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(
						middleware.Auth,
						middleware.AdminAuth,
					)
					group.Bind(
						userCtrl.NewV1(),
						adminCtrl.NewV1(),
						appVersionCtrl.NewV1(),
						completedTradesCtrl.NewV1(),
						cronTaskCtrl.NewV1(),
						copyTradingCtrl.NewV1(),
						leaderboardCtrl.NewV1(),
						walletCtrl.NewV1(),
						membershipCtrl.NewV1(),
						myTrackWalletCtrl.NewV1(),
						notificationCtrl.NewV1(),
						positionCtrl.NewV1(),
						proxyPoolCtrl.NewV1(),
						traderCtrl.NewV1(),
						traderPerformanceCtrl.NewV1(),
						traderPositionsCtrl.NewV1(),
						userAppKeyCtrl.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
