package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/admin"
	"demo/internal/controller/app_version"
	"demo/internal/controller/completed_trades"
	"demo/internal/controller/copy_trading"
	"demo/internal/controller/cron_task"
	"demo/internal/controller/leaderboard"
	"demo/internal/controller/membership"
	"demo/internal/controller/my_track_wallet"
	"demo/internal/controller/notification"
	"demo/internal/controller/position"
	"demo/internal/controller/proxy_pool"
	"demo/internal/controller/trader"
	"demo/internal/controller/trader_performance"
	"demo/internal/controller/trader_positions"
	"demo/internal/controller/user"
	"demo/internal/controller/user_app_key"
	"demo/internal/controller/wallet"
	"demo/internal/controller/ws"
	"demo/internal/initialization"
	"demo/internal/middleware"
	proxyPool "demo/internal/proxy_pool"
	"demo/internal/service"
	"demo/internal/subscriber"

	_ "demo/internal/cron_jobs"
	_ "demo/internal/logic/admin"
	_ "demo/internal/logic/app_version"
	_ "demo/internal/logic/completed_trades"
	_ "demo/internal/logic/copy_trading"
	_ "demo/internal/logic/cron_task"
	_ "demo/internal/logic/email"
	_ "demo/internal/logic/leaderboard"
	_ "demo/internal/logic/membership"
	_ "demo/internal/logic/my_track_wallet"
	_ "demo/internal/logic/notification"
	_ "demo/internal/logic/position"
	_ "demo/internal/logic/proxy_pool"
	_ "demo/internal/logic/trader"
	_ "demo/internal/logic/trader_performance"
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
				group.Bind(ws.New())
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
						user.NewV1(),
						admin.NewV1(),
						app_version.NewV1(),
						completed_trades.NewV1(),
						cron_task.NewV1(),
						copy_trading.NewV1(),
						leaderboard.NewV1(),
						wallet.NewV1(),
						membership.NewV1(),
						my_track_wallet.NewV1(),
						notification.NewV1(),
						position.NewV1(),
						proxy_pool.NewV1(),
						trader.NewV1(),
						trader_performance.NewV1(),
						trader_positions.NewV1(),
						user_app_key.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
