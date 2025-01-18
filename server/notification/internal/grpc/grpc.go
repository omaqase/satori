package grpc

import (
	"context"
	"github.com/omaqase/satori/notification/internal/config"
	"github.com/omaqase/satori/notification/internal/mailer"
	protobuf "github.com/omaqase/satori/notification/protobuf/gen"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	Listener net.Listener
	Server   *grpc.Server
}

func NewGRPCServer(lc fx.Lifecycle, mailer *mailer.Mailer, config config.Config) (*Server, error) {
	listener, err := net.Listen("tcp", ""+config.App.Port)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	reflection.Register(server)
	protobuf.RegisterNotificationServiceServer(server, mailer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error

			go func() {
				if err = server.Serve(listener); err != nil {
					log.Fatalf("failed to start grpc server: %v", err)
				}
			}()

			return err
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
