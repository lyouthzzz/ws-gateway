package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type (
	HTTPOption func(options *httpOptions)
	GRPCOption func(options *gRPCOptions)
)

type (
	httpOptions struct {
		ms      []middleware.Middleware
		srvOpts []http.ServerOption
		logger  log.Logger
	}
	gRPCOptions struct {
		ms      []middleware.Middleware
		srvOpts []grpc.ServerOption
		logger  log.Logger
	}
)

func HTTPOptionMiddleware(ms ...middleware.Middleware) HTTPOption {
	return func(options *httpOptions) { options.ms = ms }
}

func HTTPOptionServerOptions(srvOpts ...http.ServerOption) HTTPOption {
	return func(options *httpOptions) { options.srvOpts = srvOpts }
}

func HTTPOptionLogger(logger log.Logger) HTTPOption {
	return func(options *httpOptions) { options.logger = logger }
}

func GRPCOptionMiddleware(ms ...middleware.Middleware) GRPCOption {
	return func(options *gRPCOptions) { options.ms = ms }
}

func GRPCOptionServerOptions(srvOpts ...grpc.ServerOption) GRPCOption {
	return func(options *gRPCOptions) { options.srvOpts = srvOpts }
}

func GRPCOptionLogger(logger log.Logger) GRPCOption {
	return func(options *gRPCOptions) { options.logger = logger }
}
