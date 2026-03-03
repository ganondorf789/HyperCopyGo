// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notification

import (
	"context"

	"demo/api/notification/v1"
)

type INotificationV1 interface {
	NotificationSend(ctx context.Context, req *v1.NotificationSendReq) (res *v1.NotificationSendRes, err error)
	NotificationUpdate(ctx context.Context, req *v1.NotificationUpdateReq) (res *v1.NotificationUpdateRes, err error)
	NotificationDelete(ctx context.Context, req *v1.NotificationDeleteReq) (res *v1.NotificationDeleteRes, err error)
	NotificationAdminList(ctx context.Context, req *v1.NotificationAdminListReq) (res *v1.NotificationAdminListRes, err error)
	NotificationSummary(ctx context.Context, req *v1.NotificationSummaryReq) (res *v1.NotificationSummaryRes, err error)
	NotificationList(ctx context.Context, req *v1.NotificationListReq) (res *v1.NotificationListRes, err error)
	NotificationRead(ctx context.Context, req *v1.NotificationReadReq) (res *v1.NotificationReadRes, err error)
	NotificationReadAll(ctx context.Context, req *v1.NotificationReadAllReq) (res *v1.NotificationReadAllRes, err error)
}
