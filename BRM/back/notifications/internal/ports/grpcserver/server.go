package grpcserver

import (
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"notifications/internal/app"
	"notifications/internal/ports/grpcserver/pb"
	"notifications/pkg/logger"
)

//go:generate protoc pb/service.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type Server struct {
	App app.App
	pb.NotificationsServiceServer
}

func New(a app.App, logs logger.Logger) *grpc.Server {
	chain := grpcmiddleware.ChainUnaryServer(
		panicInterceptor(logs),
		loggerInterceptor(logs))

	s := grpc.NewServer(grpc.UnaryInterceptor(chain))
	pb.RegisterNotificationsServiceServer(s, &Server{
		App:                        a,
		NotificationsServiceServer: nil,
	})
	return s
}
