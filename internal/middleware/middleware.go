package middleware

import (
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"

	"demo/global"
	"demo/internal/consts"
)

type JwtClaims struct {
	UserId   int64  `json:"userId"`
	UserType string `json:"userType"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(userId int64, userType string) (string, int64, error) {
	cfg := global.Config.Jwt
	expire := cfg.Expire
	if expire == 0 {
		expire = 7200
	}

	now := time.Now()
	expireAt := now.Add(time.Duration(expire) * time.Second)

	claims := JwtClaims{
		UserId:   userId,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "HyperCopyGo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", 0, err
	}
	return tokenStr, expireAt.Unix(), nil
}

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
			claims, err := parseToken(tokenStr)
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

func extractToken(r *ghttp.Request) string {
	auth := r.GetHeader("Authorization")
	if auth != "" && strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return r.GetHeader("token")
}

func parseToken(tokenStr string) (*JwtClaims, error) {
	secret := global.Config.Jwt.Secret
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
