package server

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
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

func (x *Config) BuildHTTPServer(opts ...HTTPOption) *http.Server {
	options := &httpOptions{}
	for _, opt := range opts {
		opt(options)
	}
	timeout, err := time.ParseDuration(x.Timeout)
	if err != nil {
		panic(errors.WithMessage(err, "HTTP server invalid timeout config"))
	}
	// 默认开启
	ms := []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(options.logger),
	}
	if len(options.ms) > 0 {
		ms = append(ms, options.ms...)
	}
	srvOpts := []http.ServerOption{
		http.Address(x.Addr),
		http.Timeout(timeout),
		http.Middleware(ms...),
	}
	if len(options.srvOpts) > 0 {
		srvOpts = append(srvOpts, options.srvOpts...)
	}

	srv := http.NewServer(srvOpts...)
	return srv
}

func (x *Config) BuildGRPCServer(opts ...GRPCOption) *grpc.Server {
	options := &gRPCOptions{}
	for _, opt := range opts {
		opt(options)
	}
	timeout, err := time.ParseDuration(x.Timeout)
	if err != nil {
		panic(errors.WithMessage(err, "GRPC server invalid timeout config"))
	}
	// 默认开启
	ms := []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(options.logger),
	}
	if len(options.ms) > 0 {
		ms = append(ms, options.ms...)
	}
	srvOpts := []grpc.ServerOption{
		grpc.Address(x.Addr),
		grpc.Timeout(timeout),
		grpc.Middleware(options.ms...),
	}
	if len(options.srvOpts) > 0 {
		srvOpts = append(srvOpts, options.srvOpts...)
	}

	srv := grpc.NewServer(srvOpts...)
	return srv
}

func DefaultConfig() *Config {
	return &Config{Addr: ":8080", Timeout: "60s"}
}

func (x *Config) WithAddr(addr string) *Config {
	x.Addr = addr
	return x
}

func (x *Config) WithTimeout(d string) *Config {
	x.Timeout = d
	return x
}
