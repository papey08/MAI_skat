package grpcserver

import (
	"brm-ads/internal/app"
	"brm-ads/internal/ports/grpcserver/pb"
	"brm-ads/pkg/logger"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

//go:generate protoc pb/service.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type Server struct {
	App app.App
	pb.AdsServiceServer
}

func New(a app.App, logs logger.Logger) *grpc.Server {
	chain := grpcmiddleware.ChainUnaryServer(
		panicInterceptor(logs),
		loggerInterceptor(logs))

	s := grpc.NewServer(grpc.UnaryInterceptor(chain))
	pb.RegisterAdsServiceServer(s, &Server{
		App:              a,
		AdsServiceServer: nil,
	})
	return s
}
