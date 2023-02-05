// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/config"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/server"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/service"
)

func initApp(*config.Server, *config.Registry, log.Logger) (*kratos.App, error) {
	panic(wire.Build(service.ProviderSet, server.ProviderSet, newApp))
}
