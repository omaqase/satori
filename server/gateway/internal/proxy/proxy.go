package proxy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omaqase/satori/gateway/internal/config"
	"github.com/omaqase/satori/gateway/internal/notification/protobuf"
	"google.golang.org/grpc"
)

func SetupProxy(config *config.Config) error {
	gatewayMux := runtime.NewServeMux()

	err := SetupGRPCConnection(gatewayMux, config)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1", gatewayMux)

	return http.ListenAndServe(config.App.Port, mux)
}

func SetupGRPCConnection(gatewayMux *runtime.ServeMux, config *config.Config) error {
	grpcNotificationConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%s", config.NotificationService.Host, config.NotificationService.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	defer func(grpcNotificationConnection *grpc.ClientConn) {
		err := grpcNotificationConnection.Close()
		if err != nil {
			return
		}
	}(grpcNotificationConnection)

	err = protobuf.RegisterNotificationServiceHandler(context.Background(), gatewayMux, grpcNotificationConnection)
	if err != nil {
		return err
	}

	return nil
}
