package main

import (
	"log"

	"github.com/kalougata/mall/cmd/wire"
	"github.com/kalougata/mall/pkg/http"
)

func main() {
	server, cleanup, err := wire.NewApp()

	if err != nil {
		log.Panic(err)
	}

	http.Run(server.AdminHTTPServer, ":8001")

	defer cleanup()
}
