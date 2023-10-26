package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kalougata/mall/cmd/wire"
	"github.com/kalougata/mall/pkg/config"
	"github.com/kalougata/mall/pkg/http"
)

func main() {
	config := config.NewConfig()

	server, cleanup, err := wire.NewApp(config)

	if err != nil {
		log.Panic(err)
	}

	http.Run(server.MallHTTPServer, fmt.Sprintf(":%s", config.GetString("server.mall.port")))

	defer cleanup()
}
