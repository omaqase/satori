package main

import (
	"github.com/omaqase/satori/gateway/internal/config"
	"github.com/omaqase/satori/gateway/internal/proxy"
	"google.golang.org/grpc/grpclog"
)

func main() {
	configs, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	if err := proxy.SetupProxy(configs); err != nil {
		grpclog.Fatal(err)
	}
}
