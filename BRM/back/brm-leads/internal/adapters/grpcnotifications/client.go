package grpcnotifications

import (
	"brm-leads/internal/adapters/grpcnotifications/pb"
	"brm-leads/internal/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type notificationsClientImpl struct {
	cli pb.NotificationsServiceClient
}

func NewNotificationsClient(ctx context.Context, addr string) (NotificationsClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &notificationsClientImpl{}, fmt.Errorf("grpc notifications client: %w", err)
	} else {
		return &notificationsClientImpl{
			cli: pb.NewNotificationsServiceClient(conn),
		}, nil
	}
}

func (n *notificationsClientImpl) CreateNewLeadNotification(ctx context.Context, leadId uint64, companyId uint64, clientCompany uint64) error {
	_, err := n.cli.CreateNotification(ctx, &pb.CreateNotificationRequest{
		Notification: &pb.Notification{
			CompanyId: companyId,
			Type:      "new_lead",
			NewLead: &pb.NewLeadInfo{
				LeadId:        leadId,
				ClientCompany: clientCompany,
			},
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return model.ErrNotificationsError
		default:
			return model.ErrNotificationsError
		}
	}
	return nil
}

func (n *notificationsClientImpl) CreateCloseLeadNotification(ctx context.Context, adId uint64, producerCompany uint64, consumerCompany uint64) error {
	_, err := n.cli.CreateNotification(ctx, &pb.CreateNotificationRequest{
		Notification: &pb.Notification{
			CompanyId: consumerCompany,
			Type:      "closed_lead",
			ClosedLead: &pb.ClosedLeadInfo{
				AdId:            adId,
				ProducerCompany: producerCompany,
			},
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return model.ErrNotificationsError
		default:
			return model.ErrNotificationsError
		}
	}
	return nil
}
