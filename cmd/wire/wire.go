//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	adminv1 "github.com/kalougata/mall/api/v1/admin"
	mallv1 "github.com/kalougata/mall/api/v1/mall"
	adminctrl "github.com/kalougata/mall/controller/admin"
	"github.com/kalougata/mall/pkg/app"
	"github.com/kalougata/mall/pkg/data"
	adminrepo "github.com/kalougata/mall/repo/admin"
	adminrouter "github.com/kalougata/mall/router/admin"
	mallrouter "github.com/kalougata/mall/router/mall"
	adminsrv "github.com/kalougata/mall/service/admin"
	"github.com/spf13/viper"
)

var AdminProvider = wire.NewSet(
	adminrepo.NewUmsAdminRepo,
	adminsrv.NewUmsAdminService,
	adminctrl.NewUmsAdminController,
	adminv1.NewAdminAPIRouter,
	adminrouter.NewAdminHTTPServer,
)

var MallProvider = wire.NewSet(
	mallv1.NewMallAPIRouter,
	mallrouter.NewMallHTTPServer,
)

func NewApp(config *viper.Viper) (*app.Server, func(), error) {
	panic(wire.Build(
		data.NewData,
		AdminProvider,
		MallProvider,
		app.NewServer,
	))
}
