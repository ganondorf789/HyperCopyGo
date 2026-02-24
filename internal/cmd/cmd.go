package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	v1admin "demo/api/admin/v1"
	v1user "demo/api/user/v1"
	"demo/internal/consts"
	"demo/internal/middleware"
	"demo/internal/model"
	"demo/internal/service"

	// 注册 logic 实现
	_ "demo/internal/logic/admin"
	_ "demo/internal/logic/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// ===== 用户公开接口 =====
				group.POST("/user/register", func(r *ghttp.Request) {
					var req v1user.UserRegisterReq
					if err := r.Parse(&req); err != nil {
						r.Response.WriteJsonExit(g.Map{"code": 400, "message": err.Error()})
						return
					}
					if err := service.User().Register(r.Context(), model.UserRegisterInput{
						Username: req.Username,
						Password: req.Password,
						Nickname: req.Nickname,
					}); err != nil {
						r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
						return
					}
					r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok"})
				})

				group.POST("/user/login", func(r *ghttp.Request) {
					var req v1user.UserLoginReq
					if err := r.Parse(&req); err != nil {
						r.Response.WriteJsonExit(g.Map{"code": 400, "message": err.Error()})
						return
					}
					out, err := service.User().Login(r.Context(), model.UserLoginInput{
						Username: req.Username,
						Password: req.Password,
					})
					if err != nil {
						r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
						return
					}
					r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok", "data": out})
				})

				// ===== 管理员公开接口 =====
				group.POST("/admin/login", func(r *ghttp.Request) {
					var req v1admin.AdminLoginReq
					if err := r.Parse(&req); err != nil {
						r.Response.WriteJsonExit(g.Map{"code": 400, "message": err.Error()})
						return
					}
					out, err := service.Admin().Login(r.Context(), model.AdminLoginInput{
						Username: req.Username,
						Password: req.Password,
					})
					if err != nil {
						r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
						return
					}
					r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok", "data": out})
				})

				// ===== 用户鉴权接口 =====
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)

					group.GET("/profile", func(r *ghttp.Request) {
						userId := r.GetCtxVar(consts.CtxUserIdKey).Int64()
						out, err := service.User().Profile(r.Context(), userId)
						if err != nil {
							r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
							return
						}
						r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok", "data": out})
					})
				})

				// ===== 管理后台鉴权接口 =====
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.AdminAuth)

					group.GET("/profile", func(r *ghttp.Request) {
						adminId := r.GetCtxVar(consts.CtxUserIdKey).Int64()
						out, err := service.Admin().Profile(r.Context(), adminId)
						if err != nil {
							r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
							return
						}
						r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok", "data": out})
					})

					group.GET("/users", func(r *ghttp.Request) {
						out, err := service.Admin().UserList(r.Context(), model.AdminUserListInput{
							Page:     r.GetQuery("page", 1).Int(),
							PageSize: r.GetQuery("pageSize", 20).Int(),
							Status:   r.GetQuery("status", -1).Int(),
						})
						if err != nil {
							r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
							return
						}
						r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok", "data": out})
					})

					group.PUT("/users/:id/status", func(r *ghttp.Request) {
						id := r.Get("id").Int64()
						status := r.Get("status").Int()
						if err := service.Admin().UserSetStatus(r.Context(), model.AdminUserStatusInput{
							Id: id, Status: status,
						}); err != nil {
							r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
							return
						}
						r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok"})
					})

					group.DELETE("/users/:id", func(r *ghttp.Request) {
						id := r.Get("id").Int64()
						if err := service.Admin().UserDelete(r.Context(), id); err != nil {
							r.Response.WriteJsonExit(g.Map{"code": 500, "message": err.Error()})
							return
						}
						r.Response.WriteJsonExit(g.Map{"code": 0, "message": "ok"})
					})
				})
			})

			s.Run()
			return nil
		},
	}
)
