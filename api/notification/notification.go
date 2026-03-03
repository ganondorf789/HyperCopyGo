// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notification

import (
	"context"

	"demo/api/notification/v1"
)

type INotificationV1 interface {
	NotificationSummary(ctx context.Context, req *v1.NotificationSummaryReq) (res *v1.NotificationSummaryRes, err error)
	NotificationList(ctx context.Context, req *v1.NotificationListReq) (res *v1.NotificationListRes, err error)
	NotificationRead(ctx context.Context, req *v1.NotificationReadReq) (res *v1.NotificationReadRes, err error)
	NotificationReadAll(ctx context.Context, req *v1.NotificationReadAllReq) (res *v1.NotificationReadAllRes, err error)
}
