//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	adminv1 "github.com/kalougata/mall/api/v1/admin"
	mallv1 "github.com/kalougata/mall/api/v1/mall"
	"github.com/kalougata/mall/pkg/app"
	adminrouter "github.com/kalougata/mall/router/admin"
	mallrouter "github.com/kalougata/mall/router/mall"
	"github.com/spf13/viper"
)

func NewApp(config *viper.Viper) (*app.Server, func(), error) {
	panic(wire.Build(
		adminv1.NewAdminAPIRouter,
		mallv1.NewMallAPIRouter,
		adminrouter.NewAdminHTTPServer,
		mallrouter.NewMallHTTPServer,
		app.NewServer,
	))
}
