package main

import (
	"context"
	"flag"
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/gateway"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/upstream"
	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	wsAPIEndpoint     = flag.String("ws_api_endpoint", "127.0.0.1:8081", "ws api gRPC address")
	httpServerAddress = flag.String("http_server_addr", ":8080", "http server listen address")
)

func main() {
	flag.Parse()

	conn, err := grpc.DialContext(context.Background(), *wsAPIEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(errors.WithStack(err))
	}
	exc := exchange.NewExchangeServiceClient(conn)

	up, err := upstream.NewGRPCStreamingUpstream(
		upstream.GRPCStreamingExchangeClient(exc),
	)
	if err != nil {
		panic(errors.WithStack(err))
	}

	websocketGateway := gateway.NewWebsocketGateway(
		gateway.WebsocketGatewayOptionUpstream(up),
	)

	http.HandleFunc("/gateway/ws", websocketGateway.WebsocketConnectHandler())

	log.Println("HTTP server serve " + *httpServerAddress)

	panic(http.ListenAndServe(*httpServerAddress, nil))
}
