package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

//go:generate protoc -I. --go_out=paths=source_relative:. config.proto

func NewConfig(config config.Config, key string) (c *Config) {
	var err error
	if err = config.Value(key).Scan(&c); err != nil {
		return
	}
	err = mergo.Merge(c, DefaultConfig())
	return
}

func (x *Config) BuildHTTPClient(ctx context.Context, opts ...HTTPOption) (*http.Client, error) {
	options := &httpOptions{}
	for _, opt := range opts {
		opt(options)
	}
	timeout, err := time.ParseDuration(x.Timeout)
	if err != nil {
		panic(errors.WithMessage(err, "HTTP client invalid timeout config"))
	}
	// 默认开启
	ms := []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(options.logger),
	}
	if len(options.ms) > 0 {
		ms = append(ms, options.ms...)
	}
	clientOpts := []http.ClientOption{
		http.WithEndpoint(x.Addr),
		http.WithTimeout(timeout),
		http.WithMiddleware(ms...),
	}
	if len(options.clientOpts) > 0 {
		clientOpts = append(clientOpts, options.clientOpts...)
	}

	return http.NewClient(ctx, clientOpts...)
}

func (x *Config) BuildGRPCClient(ctx context.Context, opts ...GRPCOption) (*gRPC.ClientConn, error) {
	options := &gRPCOptions{}
	for _, opt := range opts {
		opt(options)
	}
	timeout, err := time.ParseDuration(x.Timeout)
	if err != nil {
		return nil, errors.WithMessage(err, "GRPC client invalid timeout config")
	}
	// 默认开启
	ms := []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(options.logger),
	}
	if len(options.ms) > 0 {
		ms = append(ms, options.ms...)
	}
	clientOpts := []grpc.ClientOption{
		grpc.WithEndpoint(x.Addr),
		grpc.WithTimeout(timeout),
		grpc.WithMiddleware(options.ms...),
		grpc.WithOptions(gRPC.WithTransportCredentials(insecure.NewCredentials())),
	}
	if len(options.clientOpts) > 0 {
		clientOpts = append(clientOpts, options.clientOpts...)
	}

	return grpc.DialInsecure(ctx, clientOpts...)
}

func DefaultConfig() *Config {
	return &Config{Timeout: "5s"}
}

func (x *Config) WithAddr(addr string) *Config {
	x.Addr = addr
	return x
}

func (x *Config) WithTimeout(d string) *Config {
	x.Timeout = d
	return x
}
