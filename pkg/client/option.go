package client

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
		ms         []middleware.Middleware
		clientOpts []http.ClientOption
		logger     log.Logger
	}
	gRPCOptions struct {
		ms         []middleware.Middleware
		clientOpts []grpc.ClientOption
		logger     log.Logger
	}
)

func HTTPOptionMiddleware(ms ...middleware.Middleware) HTTPOption {
	return func(options *httpOptions) { options.ms = ms }
}

func HTTPOptionClientOptions(srvOpts ...http.ClientOption) HTTPOption {
	return func(options *httpOptions) { options.clientOpts = srvOpts }
}

func HTTPOptionLogger(logger log.Logger) HTTPOption {
	return func(options *httpOptions) { options.logger = logger }
}

func GRPCOptionMiddleware(ms ...middleware.Middleware) GRPCOption {
	return func(options *gRPCOptions) { options.ms = ms }
}

func GRPCOptionClientOptions(srvOpts ...grpc.ClientOption) GRPCOption {
	return func(options *gRPCOptions) { options.clientOpts = srvOpts }
}

func GRPCOptionLogger(logger log.Logger) GRPCOption {
	return func(options *gRPCOptions) { options.logger = logger }
}
