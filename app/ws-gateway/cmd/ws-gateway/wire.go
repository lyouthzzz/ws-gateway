// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/config"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/gateway"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/server"
)

func initApp(*config.Server, *etcd.Registry, *gateway.WebsocketGateway, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
