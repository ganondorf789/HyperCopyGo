package model

// ==================== 配置 ====================

type Config struct {
	Jwt      JwtConfig
	SendGrid SendGridConfig
}

type SendGridConfig struct {
	ApiKey    string `json:"apiKey"`
	FromEmail string `json:"fromEmail"`
	FromName  string `json:"fromName"`
}

type JwtConfig struct {
	Secret string `json:"secret"`
	Expire int64  `json:"expire"`
}

// ==================== 管理员 ====================
