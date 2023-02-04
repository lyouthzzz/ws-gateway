package main

import (
	"flag"
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/service"
	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/grpc"
	channelzservice "google.golang.org/grpc/channelz/service"
	"log"
	"net"
)

var (
	gRPCServerAddress = flag.String("grpc_server_address", ":8081", "")
)

func main() {
	grpcServer := grpc.NewServer()

	exchange.RegisterExchangeServiceServer(grpcServer, service.NewExchangeService())
	channelzservice.RegisterChannelzServiceToServer(grpcServer)

	lis, err := net.Listen("tcp", *gRPCServerAddress)
	if err != nil {
		panic(errors.WithStack(err))
	}
	log.Println("gRPC server serve " + *gRPCServerAddress)
	if err := grpcServer.Serve(lis); err != nil {
		panic(errors.WithStack(err))
	}
}
