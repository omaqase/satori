package proxy

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omaqase/satori/gateway/internal/config"
	"github.com/omaqase/satori/gateway/internal/notification/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50054", "gRPC server endpoint")
)

func SetupProxy(config config.Config) error {
	//gatewayMux := runtime.NewServeMux()

	//grpcNotificationConnection, err := grpc.Dial(
	//	"localhost:50054",
	//	grpc.WithInsecure(),
	//)
	//log.Println("q")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//defer grpcNotificationConnection.Close()
	//
	//err = protobuf.RegisterNotificationServiceHandler(context.Background(), gatewayMux, grpcNotificationConnection)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	//mux := http.NewServeMux()
	//mux.Handle("/api/v1", gatewayMux)
	//mux.HandleFunc("/", helloworld)

	//swaggerPath := "swagger/notification.swagger.json" // Путь к Swagger-файлу
	//mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, swaggerPath)
	//})
	//
	//fmt.Println(config)
	//http.ListenAndServe(":"+config.App.Port, mux)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := protobuf.RegisterNotificationServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
