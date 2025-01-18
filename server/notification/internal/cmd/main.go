package main

import (
	"github.com/omaqase/satori/notification/internal/config"
	"github.com/omaqase/satori/notification/internal/grpc"
	"github.com/omaqase/satori/notification/internal/mailer"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			mailer.NewResendClient,
			mailer.NewMailer,
			grpc.NewGRPCServer,
		),
		fx.Invoke(func(config *grpc.Server) {}),
	)

	app.Run()
}
