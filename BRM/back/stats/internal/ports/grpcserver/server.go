package grpcserver

import (
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"stats/internal/app"
	"stats/internal/ports/grpcserver/pb"
	"stats/pkg/logger"
)

//go:generate protoc pb/service.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type Server struct {
	a app.App
	pb.StatsServiceServer
}

func New(a app.App, logs logger.Logger) *grpc.Server {
	chain := grpcmiddleware.ChainUnaryServer(
		panicInterceptor(logs),
		loggerInterceptor(logs))

	s := grpc.NewServer(grpc.UnaryInterceptor(chain))
	pb.RegisterStatsServiceServer(s, &Server{
		a:                  a,
		StatsServiceServer: nil,
	})
	return s
}
