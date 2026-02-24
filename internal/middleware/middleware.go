package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"demo/internal/consts"
)

type JwtClaims struct {
	UserId   int64  `json:"userId"`
	UserType string `json:"userType"` // "user" or "admin"
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(userId int64, userType string) (string, int64, error) {
	secret := g.Cfg().MustGet(nil, "jwt.secret").String()
	expire := g.Cfg().MustGet(nil, "jwt.expire").Int64()
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
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenStr, expireAt.Unix(), nil
}

// Auth JWT 鉴权中间件
func Auth(r *ghttp.Request) {
	tokenStr := extractToken(r)
	if tokenStr == "" {
		r.Response.WriteJsonExit(g.Map{
			"code":    http.StatusUnauthorized,
			"message": "未登录或token已过期",
		})
		return
	}

	claims, err := parseToken(tokenStr)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    http.StatusUnauthorized,
			"message": "token无效",
		})
		return
	}

	r.SetCtxVar(consts.CtxUserIdKey, claims.UserId)
	r.SetCtxVar(consts.CtxUserTypeKey, claims.UserType)
	r.Middleware.Next()
}

// AdminAuth 后台管理鉴权中间件（要求 userType == admin）
func AdminAuth(r *ghttp.Request) {
	tokenStr := extractToken(r)
	if tokenStr == "" {
		r.Response.WriteJsonExit(g.Map{
			"code":    http.StatusUnauthorized,
			"message": "未登录或token已过期",
		})
		return
	}

	claims, err := parseToken(tokenStr)
	if err != nil || claims.UserType != consts.UserTypeAdmin {
		r.Response.WriteJsonExit(g.Map{
			"code":    http.StatusForbidden,
			"message": "无权限访问",
		})
		return
	}

	r.SetCtxVar(consts.CtxUserIdKey, claims.UserId)
	r.SetCtxVar(consts.CtxUserTypeKey, claims.UserType)
	r.Middleware.Next()
}

func extractToken(r *ghttp.Request) string {
	auth := r.GetHeader("Authorization")
	if auth != "" && strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return r.Get("token").String()
}

func parseToken(tokenStr string) (*JwtClaims, error) {
	secret := g.Cfg().MustGet(nil, "jwt.secret").String()
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
