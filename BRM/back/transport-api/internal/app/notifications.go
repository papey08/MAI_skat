package app

import (
	"context"
	"transport-api/internal/model/notifications"
)

func (a *appImpl) GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]notifications.Notification, uint, error) {
	return a.notifications.GetNotifications(ctx, companyId, limit, offset, onlyNotViewed)
}

func (a *appImpl) GetNotification(ctx context.Context, companyId uint64, notificationId uint64) (notifications.Notification, error) {
	return a.notifications.GetNotification(ctx, companyId, notificationId)
}

func (a *appImpl) SubmitClosedLead(ctx context.Context, companyId uint64, notificationId uint64, submit bool) error {
	return a.notifications.SubmitClosedLead(ctx, companyId, notificationId, submit)
}
