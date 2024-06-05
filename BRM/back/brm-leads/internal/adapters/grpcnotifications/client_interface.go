package grpcnotifications

import "context"

//go:generate protoc pb/notifications_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type NotificationsClient interface {
	CreateNewLeadNotification(ctx context.Context, leadId uint64, companyId uint64, clientCompany uint64) error
	CreateCloseLeadNotification(ctx context.Context, adId uint64, producerCompany uint64, consumerCompany uint64) error
}
