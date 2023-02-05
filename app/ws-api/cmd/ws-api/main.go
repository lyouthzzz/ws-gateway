package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	appconfig "github.com/lyouthzzz/ws-gateway/app/ws-api/internal/config"
	"github.com/lyouthzzz/ws-gateway/pkg/env"
	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
)

var (
	configPath = flag.String("config", "", "config file path of project")
)

func newApp(logger log.Logger, svrs []transport.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(env.AppName),
		kratos.Version(env.AppVersion),
		kratos.Logger(logger),
		kratos.Server(svrs...),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()

	c := config.New(config.WithSource(file.NewSource(*configPath)))
	if err := c.Load(); err != nil {
		panic(errors.WithStack(err))
	}
	var bc appconfig.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(errors.WithStack(err))
	}
	app, err := initApp(bc.Server, bc.Registry, log.DefaultLogger)
	if err != nil {
		panic(errors.WithStack(err))
	}
	if err := app.Run(); err != nil {
		panic(errors.WithStack(err))
	}
}
