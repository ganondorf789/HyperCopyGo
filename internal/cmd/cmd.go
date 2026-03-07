package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/admin"
	"demo/internal/controller/app_version"
	"demo/internal/controller/copy_trade_config"
	"demo/internal/controller/copy_trading"
	"demo/internal/controller/cron_task"
	"demo/internal/controller/leaderboard"
	"demo/internal/controller/membership"
	"demo/internal/controller/my_track_wallet"
	"demo/internal/controller/notification"
	"demo/internal/controller/position"
	"demo/internal/controller/proxy_pool"
	"demo/internal/controller/trade"
	"demo/internal/controller/trader"
	"demo/internal/controller/trader_performance"
	"demo/internal/controller/trader_positions"
	"demo/internal/controller/user"
	"demo/internal/controller/system_setting"
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
	_ "demo/internal/logic/copy_trade_config"
	_ "demo/internal/logic/copy_trading"
	_ "demo/internal/logic/cron_task"
	_ "demo/internal/logic/email"
	_ "demo/internal/logic/leaderboard"
	_ "demo/internal/logic/membership"
	_ "demo/internal/logic/my_track_wallet"
	_ "demo/internal/logic/notification"
	_ "demo/internal/logic/position"
	_ "demo/internal/logic/proxy_pool"
	_ "demo/internal/logic/trade"
	_ "demo/internal/logic/trader"
	_ "demo/internal/logic/trader_performance"
	_ "demo/internal/logic/trader_positions"
	_ "demo/internal/logic/user"
	_ "demo/internal/logic/system_setting"
	_ "demo/internal/logic/user_app_key"
	_ "demo/internal/logic/wallet"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 鍒濆鍖栭厤缃?
			if err = initialization.InitConfig(); err != nil {
				return err
			}

			// 鍒濆鍖栦唬鐞嗘睜锛堜粠鏁版嵁搴撳姞杞斤級
			if err = proxyPool.Reload(); err != nil {
				return err
			}

			// 浠庢暟鎹簱鍔犺浇骞跺惎鍔ㄥ畾鏃朵换鍔?
			service.CronTask().StartAll(ctx)

			// 鍚姩 Redis 璁㈤槄锛坣ew_positions / market_alert锛?
			subscriber.Start(ctx)

			s := g.Server()

			// WebSocket 璺敱锛堜笉浣跨敤 MiddlewareHandlerResponse锛岄伩鍏嶅共鎵?WebSocket 鍗囩骇锛?
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.CORS)
				group.Bind(ws.New())
			})

			// REST API 璺敱
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(
						middleware.CORS,
						middleware.Auth,
						middleware.AdminAuth,
					)
					group.Bind(
						user.NewV1(),
						admin.NewV1(),
						app_version.NewV1(),
					cron_task.NewV1(),
						copy_trade_config.NewV1(),
					copy_trading.NewV1(),
						leaderboard.NewV1(),
						wallet.NewV1(),
						membership.NewV1(),
						my_track_wallet.NewV1(),
						notification.NewV1(),
						position.NewV1(),
					proxy_pool.NewV1(),
					system_setting.NewV1(),
					trader.NewV1(),
						trader_performance.NewV1(),
						trader_positions.NewV1(),
						trade.NewV1(),
						user_app_key.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)

