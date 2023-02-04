package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	appconfig "github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/config"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/gateway"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/upstream"
	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
	_ "net/http/pprof"
	"os"
)

var (
	appName    = "ws-gateway"
	appVersion = "v0.0.1"
)

func init() {
	// 应用名称 发布平台通过环境变量注入
	appName = os.Getenv("APP_NAME")
	// 应用版本 发布平台通过环境变量注入 通常是发布的镜像名称
	appVersion = os.Getenv("APP_VERSION")
}

var (
	configPath = flag.String("config", "", "config file path of project")
)

func newApp(logger log.Logger, svrs []transport.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(appName),
		kratos.Version(appVersion),
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

	wsAPIClient, err := bc.Client.WsAPI.BuildGRPCClient(context.Background())
	if err != nil {
		panic(errors.WithStack(err))
	}
	exc := exchange.NewExchangeServiceClient(wsAPIClient)
	up, err := upstream.NewGRPCStreamingUpstream(
		upstream.GRPCStreamingExchangeClient(exc),
	)
	if err != nil {
		panic(errors.WithStack(err))
	}

	websocketGateway := gateway.NewWebsocketGateway(
		gateway.WebsocketGatewayOptionUpstream(up),
	)

	app, err := initApp(bc.Server, bc.Registry, websocketGateway, log.DefaultLogger)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
