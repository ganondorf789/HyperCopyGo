package service

import "context"

type IEmail interface {
	SendVerifyCode(ctx context.Context, email string, code string) error
}

var localEmail IEmail

func Email() IEmail {
	return localEmail
}

func RegisterEmail(s IEmail) {
	localEmail = s
}
