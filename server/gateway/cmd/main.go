package main

import (
	"github.com/omaqase/satori/gateway/internal/config"
	"github.com/omaqase/satori/gateway/internal/proxy"
)

func main() {
	configs, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	proxy.SetupProxy(configs)
}
