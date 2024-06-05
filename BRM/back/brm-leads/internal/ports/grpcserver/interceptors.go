package grpcserver

import (
	"brm-leads/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func loggerInterceptor(logs logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		logs.Info(logger.Fields{
			"Method": info.FullMethod,
		}, "got request")
		return handler(ctx, req)
	}
}

func panicInterceptor(logs logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error(logger.Fields{
					"Method": info.FullMethod,
				}, fmt.Sprintf("panic: %v", r))
			}
		}()
		return handler(ctx, req)
	}
}
