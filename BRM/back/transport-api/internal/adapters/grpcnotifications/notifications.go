package grpcnotifications

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transport-api/internal/adapters/grpcnotifications/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/notifications"
)

func (n *notificationsClientImpl) GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]notifications.Notification, uint, error) {
	resp, err := n.cli.GetNotifications(ctx, &pb.GetNotificationsRequest{
		CompanyId:     companyId,
		Limit:         uint64(limit),
		Offset:        uint64(offset),
		OnlyNotViewed: onlyNotViewed,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return []notifications.Notification{}, 0, model.ErrNotificationsError
		default:
			return []notifications.Notification{}, 0, model.ErrNotificationsError
		}
	}

	notificationsList := make([]notifications.Notification, len(resp.List))
	for i, notification := range resp.List {
		notificationsList[i] = respToNotification(notification)
	}
	return notificationsList, uint(resp.Amount), nil
}

func (n *notificationsClientImpl) GetNotification(ctx context.Context, companyId uint64, notificationId uint64) (notifications.Notification, error) {
	resp, err := n.cli.GetNotification(ctx, &pb.GetNotificationRequest{
		CompanyId:      companyId,
		NotificationId: notificationId,
	})
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return notifications.Notification{}, model.ErrNotificationNotExists
		case codes.PermissionDenied:
			return notifications.Notification{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return notifications.Notification{}, model.ErrNotificationsError
		default:
			return notifications.Notification{}, model.ErrNotificationsError
		}
	}
	return respToNotification(resp.Notification), nil
}

func (n *notificationsClientImpl) SubmitClosedLead(ctx context.Context, companyId uint64, notificationId uint64, submit bool) error {
	_, err := n.cli.SubmitClosedLead(ctx, &pb.SubmitClosedLeadRequest{
		CompanyId:      companyId,
		NotificationId: notificationId,
		Submit:         submit,
	})
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return model.ErrNotificationNotExists
		case codes.PermissionDenied:
			return model.ErrPermissionDenied
		case codes.FailedPrecondition:
			return model.ErrInvalidInput
		case codes.Canceled:
			return model.ErrNotificationAnswered
		case codes.ResourceExhausted:
			return model.ErrNotificationsError
		default:
			return model.ErrNotificationsError
		}
	}
	return nil
}

func respToNotification(notification *pb.Notification) notifications.Notification {
	if notification == nil {
		return notifications.Notification{}
	}
	var res notifications.Notification
	res.Id = notification.Id
	res.CompanyId = notification.CompanyId
	res.Type = notification.Type
	res.Date = notification.Date
	res.Viewed = notification.Viewed

	if notification.NewLead != nil {
		res.NewLead = new(notifications.NewLead)
		res.NewLead.LeadId = notification.NewLead.LeadId
		res.NewLead.ClientCompany = notification.NewLead.ClientCompany
	}

	if notification.ClosedLead != nil {
		res.ClosedLead = new(notifications.ClosedLead)
		res.ClosedLead.AdId = notification.ClosedLead.AdId
		res.ClosedLead.ProducerCompany = notification.ClosedLead.ProducerCompany
		res.ClosedLead.Answered = notification.ClosedLead.Answered
	}
	return res
}
