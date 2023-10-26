package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/kalougata/mall/cmd/wire"
	"github.com/kalougata/mall/pkg/http"
)

func main() {
	server, cleanup, err := wire.NewApp()

	if err != nil {
		log.Panic(err)
	}

	http.Run(server.MallHTTPServer, ":8000")

	defer cleanup()
}
