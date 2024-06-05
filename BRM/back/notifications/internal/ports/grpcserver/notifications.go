package grpcserver

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"notifications/internal/model"
	"notifications/internal/model/notifications"
	"notifications/internal/ports/grpcserver/pb"
)

func (s *Server) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*empty.Empty, error) {
	var notification model.Notification
	notification.CompanyId = req.Notification.CompanyId
	notification.Type = model.NotificationType(req.Notification.Type)
	switch notification.Type {
	case model.NewLead:
		notification.NewLead = new(notifications.NewLead)
		notification.NewLead.LeadId = req.Notification.NewLead.LeadId
		notification.NewLead.ClientCompany = req.Notification.NewLead.ClientCompany
	case model.ClosedLead:
		notification.ClosedLead = new(notifications.ClosedLead)
		notification.ClosedLead.AdId = req.Notification.ClosedLead.AdId
		notification.ClosedLead.ProducerCompany = req.Notification.ClosedLead.ProducerCompany
	}
	err := s.App.CreateNotification(ctx, notification)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetNotifications(ctx context.Context, req *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	notificationsList, amount, err := s.App.GetNotifications(ctx,
		req.CompanyId,
		uint(req.Limit),
		uint(req.Offset),
		req.OnlyNotViewed,
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	resp := &pb.GetNotificationsResponse{
		List:   make([]*pb.Notification, len(notificationsList)),
		Amount: uint64(amount),
	}
	for i := range notificationsList {
		resp.List[i] = &pb.Notification{
			Id:        notificationsList[i].Id,
			CompanyId: notificationsList[i].CompanyId,
			Type:      string(notificationsList[i].Type),
			Date:      notificationsList[i].Date.Unix(),
			Viewed:    notificationsList[i].Viewed,
		}
	}
	return resp, nil
}

func (s *Server) GetNotification(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	notification, err := s.App.GetNotification(ctx, req.CompanyId, req.NotificationId)
	if err != nil {
		return nil, mapErrors(err)
	}
	resp := &pb.GetNotificationResponse{
		Notification: &pb.Notification{
			Id:        notification.Id,
			CompanyId: notification.CompanyId,
			Type:      string(notification.Type),
			Date:      notification.Date.Unix(),
			Viewed:    notification.Viewed,
		},
	}
	switch notification.Type {
	case model.NewLead:
		resp.Notification.NewLead = &pb.NewLeadInfo{
			LeadId:        notification.NewLead.LeadId,
			ClientCompany: notification.NewLead.ClientCompany,
		}
	case model.ClosedLead:
		resp.Notification.ClosedLead = &pb.ClosedLeadInfo{
			AdId:            notification.ClosedLead.AdId,
			ProducerCompany: notification.ClosedLead.ProducerCompany,
			Answered:        notification.ClosedLead.Answered,
		}
	}
	return resp, nil
}

func (s *Server) SubmitClosedLead(ctx context.Context, req *pb.SubmitClosedLeadRequest) (*empty.Empty, error) {
	err := s.App.SubmitClosedLead(ctx, req.CompanyId, req.NotificationId, req.Submit)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}
