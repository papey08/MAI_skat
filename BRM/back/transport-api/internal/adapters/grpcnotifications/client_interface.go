package grpcnotifications

import (
	"context"
	"transport-api/internal/model/notifications"
)

//go:generate protoc pb/notifications_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type NotificationsClient interface {
	GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]notifications.Notification, uint, error)
	GetNotification(ctx context.Context, companyId uint64, notificationId uint64) (notifications.Notification, error)
	SubmitClosedLead(ctx context.Context, companyId uint64, notificationId uint64, submit bool) error
}
