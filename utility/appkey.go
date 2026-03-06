package utility

import (
	"context"
	"fmt"

	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/entity"
)

// ValidateAppKey 校验 app_id/app_secret，返回 userId
func ValidateAppKey(ctx context.Context, appId, appSecret string) (int64, error) {
	var appKey entity.UserAppKey
	err := dao.UserAppKey.Ctx(ctx).
		Where("app_id = ? AND app_secret = ? AND status = ?", appId, appSecret, consts.UserStatusEnabled).
		Scan(&appKey)
	if err != nil {
		return 0, err
	}
	if appKey.Id == 0 {
		return 0, fmt.Errorf("invalid app credentials")
	}
	return appKey.UserId, nil
}
