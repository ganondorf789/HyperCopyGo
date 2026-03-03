package utility

import (
	"time"

	"demo/global"

	"github.com/golang-jwt/jwt/v5"
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

// ParseToken 解析并验证 JWT token
func ParseToken(tokenStr string) (*JwtClaims, error) {
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
