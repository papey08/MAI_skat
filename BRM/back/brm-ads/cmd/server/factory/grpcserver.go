package factory

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
)

func PrepareListener() (net.Listener, error) {
	if lis, err := net.Listen(
		viper.GetString("grpc-server.network"),
		fmt.Sprintf(":%d", viper.GetInt("grpc-server.port")),
	); err != nil {
		return nil, fmt.Errorf("creating listener %w", err)
	} else {
		return lis, nil
	}
}
