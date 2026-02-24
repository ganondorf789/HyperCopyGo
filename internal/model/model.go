package model

// ==================== 配置 ====================

type Config struct {
	Jwt JwtConfig
}

type JwtConfig struct {
	Secret string `json:"secret"`
	Expire int64  `json:"expire"`
}

// ==================== 管理员 ====================
