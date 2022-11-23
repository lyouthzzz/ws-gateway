package service

import (
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
)

func NewExchangeService() *ExchangeService {
	return &ExchangeService{streams: make(map[string]exchange.ExchangeService_ExchangeMsgServer), logger: log.Default()}
}

type ExchangeService struct {
	exchange.UnimplementedExchangeServiceServer
	streams map[string]exchange.ExchangeService_ExchangeMsgServer

	logger *log.Logger
}

func (exchangeService *ExchangeService) ExchangeMsg(msgServer exchange.ExchangeService_ExchangeMsgServer) error {

	md, _ := metadata.FromIncomingContext(msgServer.Context())
	gatewayIP := md.Get("X-Gateway-IP")
	if len(gatewayIP) == 0 || gatewayIP[0] == "" {
		return errors.New("gateway ip header not found")
	}

	exchangeService.streams[gatewayIP[0]] = msgServer
	exchangeService.logger.Printf("gateway %s gRPC streaming connect\n", gatewayIP)

	defer func() {
		exchangeService.logger.Printf("gateway %s gRPC streaming closed\n", gatewayIP)
		delete(exchangeService.streams, gatewayIP[0])
	}()

	for {
		msg, err := msgServer.Recv()
		if errors.Is(err, io.EOF) {
			exchangeService.logger.Printf("gateway stream is closed\n")
			break
		}
		if err != nil {
			exchangeService.logger.Printf("recv err: %s\n", err.Error())
			break
		}
		if msg.Payload == nil {
			exchangeService.logger.Printf("recv payload must not nil\n")
			continue
		}
		switch msg.Payload.Type {
		default:
			exchangeService.logger.Println(msg)
		}
	}

	return nil
}
