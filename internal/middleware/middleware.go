package middleware

import (
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"

	"demo/internal/consts"
	"demo/utility"
)

// Auth JWT 鉴权中间件，通过 Meta 标签 login_required:"true" 要求登录
func Auth(r *ghttp.Request) {
	handler := r.GetServeHandler().Handler
	if handler.Info.Type != nil && handler.Info.Type.NumIn() == 2 {
		var objectReq = reflect.New(handler.Info.Type.In(1))
		if v := gmeta.Get(objectReq, "login_required"); !v.IsEmpty() && v.Bool() {
			tokenStr := extractToken(r)
			if tokenStr == "" {
				r.Response.WriteJsonExit(g.Map{"code": 401, "message": "未登录"})
				return
			}
			claims, err := utility.ParseToken(tokenStr)
			if err != nil {
				r.Response.WriteJsonExit(g.Map{"code": 401, "message": "token无效"})
				return
			}
			r.SetCtxVar(consts.CtxUserIdKey, claims.UserId)
			r.SetCtxVar(consts.CtxUserTypeKey, claims.UserType)
		}
	}
	r.Middleware.Next()
}

// AdminAuth 管理员权限中间件，通过 Meta 标签 admin_required:"true" 要求管理员身份
func AdminAuth(r *ghttp.Request) {
	handler := r.GetServeHandler().Handler
	if handler.Info.Type != nil && handler.Info.Type.NumIn() == 2 {
		var objectReq = reflect.New(handler.Info.Type.In(1))
		if v := gmeta.Get(objectReq, "admin_required"); !v.IsEmpty() && v.Bool() {
			userType := r.GetCtxVar(consts.CtxUserTypeKey).String()
			if userType != consts.UserTypeAdmin {
				r.Response.WriteJsonExit(g.Map{"code": 403, "message": "没有权限"})
				return
			}
		}
	}
	r.Middleware.Next()
}

// CORS 跨域中间件，允许所有域名
func CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"*"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func extractToken(r *ghttp.Request) string {
	auth := r.GetHeader("Authorization")
	if auth != "" && strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return r.GetHeader("token")
}

