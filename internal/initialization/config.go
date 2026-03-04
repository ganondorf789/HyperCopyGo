package initialization

import (
	"demo/global"
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func InitConfig() (err error) {
	var ctx = gctx.New()

	var jwtConfig model.JwtConfig
	jwtVar := g.Cfg().MustGet(ctx, "jwt")
	err = jwtVar.Scan(&jwtConfig)
	if err != nil {
		return err
	}

	var sendGridConfig model.SendGridConfig
	sendGridVar := g.Cfg().MustGet(ctx, "sendgrid")
	err = sendGridVar.Scan(&sendGridConfig)
	if err != nil {
		return err
	}

	global.Config = &model.Config{
		Jwt:      jwtConfig,
		SendGrid: sendGridConfig,
	}
	return nil
}
