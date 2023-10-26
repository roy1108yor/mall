package app

import (
	adminrouter "github.com/kalougata/mall/router/admin"
	mallrouter "github.com/kalougata/mall/router/mall"
)

type Server struct {
	AdminHTTPServer adminrouter.AdminHTTPServer
	MallHTTPServer  mallrouter.MallHTTPServer
}

func NewServer(
	AdminHTTPServer adminrouter.AdminHTTPServer,
	MallHTTPServer mallrouter.MallHTTPServer,
) *Server {
	return &Server{
		AdminHTTPServer,
		MallHTTPServer,
	}
}
