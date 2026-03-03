// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/notification/v1"
)

type (
	INotification interface {
		Summary(ctx context.Context, userId int64) (res *v1.NotificationSummaryRes, err error)
		List(ctx context.Context, userId int64, in v1.NotificationListReq) (res *v1.NotificationListRes, err error)
		Read(ctx context.Context, userId int64, ids []int64) error
		ReadAll(ctx context.Context, userId int64, category string) error
	}
)

var (
	localNotification INotification
)

func Notification() INotification {
	if localNotification == nil {
		panic("implement not found for interface INotification, forgot register?")
	}
	return localNotification
}

func RegisterNotification(i INotification) {
	localNotification = i
}
