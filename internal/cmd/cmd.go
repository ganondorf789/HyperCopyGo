package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/admin"
	"demo/internal/controller/user"
	"demo/internal/initialization"
	"demo/internal/middleware"

	_ "demo/internal/logic/admin"
	_ "demo/internal/logic/user"
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
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
