package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/lyouthzzz/ws-gateway/api/wsapi"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/config"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewServers, NewRegistry)

func NewServers(conf *config.Server, exchangeSvc *service.ExchangeService) []transport.Server {
	httpServer := conf.Http.BuildHTTPServer()
	_PromHTTP_Metrics_Handler(httpServer)

	grpcServer := conf.Grpc.BuildGRPCServer()
	wsapi.RegisterExchangeServiceServer(grpcServer, exchangeSvc)
	//channelzservice.RegisterChannelzServiceToServer(grpcServer)
	return []transport.Server{httpServer, grpcServer}
}

func NewRegistry(conf *config.Registry) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{conf.Etcd.Addr}, DialTimeout: conf.Etcd.DialTimeout.AsDuration()})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

// 注册 prometheus metrics 路由
func _PromHTTP_Metrics_Handler(srv *http.Server) {
	srv.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{EnableOpenMetrics: true}))
}
