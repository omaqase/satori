package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/omaqase/satori/notification/internal/config"
	"github.com/omaqase/satori/notification/internal/mailer"
	protobuf "github.com/omaqase/satori/notification/protobuf/gen"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Listener net.Listener
	Server   *grpc.Server
}

func NewGRPCServer(lc fx.Lifecycle, mailer *mailer.Mailer, config config.Config) (*Server, error) {
	addr := fmt.Sprintf(":%s", config.App.Port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	reflection.Register(server)
	protobuf.RegisterNotificationServiceServer(server, mailer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Serve(listener); err != nil {
					log.Printf("Failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})

	return &Server{
		Listener: listener,
		Server:   server,
	}, nil
}
