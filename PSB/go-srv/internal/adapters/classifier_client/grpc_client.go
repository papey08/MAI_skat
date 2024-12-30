package classifier_client

import (
	"context"
	"go-srv/internal/adapters/classifier_client/classifier"

	"google.golang.org/grpc"
)

//go:generate protoc classifier.proto --proto_path=. --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type ClassifierCli struct {
	cli classifier.ProtoServiceClient
}

func NewClassifierCli(ctx context.Context, addr string) (*ClassifierCli, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &ClassifierCli{
		cli: classifier.NewProtoServiceClient(conn),
	}, nil
}

func (c *ClassifierCli) Predict(ctx context.Context, text string) (string, error) {
	resp, err := c.cli.Predict(ctx, &classifier.PredictRequest{OriginalText: text})
	if err != nil {
		return "", err
	}
	return resp.Category, nil
}
