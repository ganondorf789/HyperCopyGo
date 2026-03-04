package email

import (
	"context"
	"fmt"

	"demo/global"
	"demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func init() {
	service.RegisterEmail(&sEmail{})
}

type sEmail struct{}

func (s *sEmail) SendVerifyCode(ctx context.Context, toEmail string, code string) error {
	cfg := global.Config.SendGrid
	from := mail.NewEmail(cfg.FromName, cfg.FromEmail)
	to := mail.NewEmail("", toEmail)
	subject := "Your Login Verification Code"
	htmlContent := fmt.Sprintf(
		`<div style="font-family:Arial,sans-serif;max-width:480px;margin:0 auto;padding:32px;border:1px solid #e0e0e0;border-radius:8px">
			<h2 style="color:#333">Verification Code</h2>
			<p style="font-size:16px;color:#555">Your verification code is:</p>
			<div style="font-size:32px;font-weight:bold;letter-spacing:8px;color:#1a73e8;text-align:center;padding:16px 0">%s</div>
			<p style="font-size:14px;color:#999">This code will expire in 5 minutes. Do not share it with anyone.</p>
		</div>`, code)

	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(cfg.ApiKey)
	resp, err := client.SendWithContext(ctx, message)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		g.Log().Errorf(ctx, "SendGrid error: status=%d body=%s", resp.StatusCode, resp.Body)
		return fmt.Errorf("邮件发送失败")
	}
	return nil
}
