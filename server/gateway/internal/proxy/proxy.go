package proxy

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omaqase/satori/gateway/internal/config"
	"github.com/omaqase/satori/gateway/internal/notification/protobuf"
	"google.golang.org/grpc"
)

func SetupProxy(config config.Config) {
	gatewayMux := runtime.NewServeMux()

	grpcNotificationConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", config.NotificationService.Host, config.NotificationService.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer grpcNotificationConnection.Close()

	err = protobuf.RegisterNotificationServiceHandler(context.Background(), gatewayMux, grpcNotificationConnection)
	if err != nil {
		log.Fatalf(err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1", gatewayMux)

	fmt.Println(config)
	http.ListenAndServe(":"+config.App.Port, mux)
}

func SetupGRPCConnection(gatewayMux *runtime.ServeMux, config config.Config) error {
	grpcNotificationConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", config.NotificationService.Host, config.NotificationService.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	defer grpcNotificationConnection.Close()

	err = protobuf.RegisterNotificationServiceHandler(context.Background(), gatewayMux, grpcNotificationConnection)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
