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
		Update(ctx context.Context, in v1.NotificationUpdateReq) error
		Delete(ctx context.Context, id int64) error
		AdminList(ctx context.Context, in v1.NotificationAdminListReq) (res *v1.NotificationAdminListRes, err error)
		Send(ctx context.Context, in v1.NotificationSendReq) (res *v1.NotificationSendRes, err error)
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
