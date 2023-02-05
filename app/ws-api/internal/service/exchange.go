package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/dispatcher"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"io"
)

var ProviderSet = wire.NewSet(NewExchangeService)

func NewExchangeService() *ExchangeService {
	return &ExchangeService{streams: make(map[string]exchange.ExchangeService_ExchangeMsgServer)}
}

type ExchangeService struct {
	exchange.UnimplementedExchangeServiceServer
	streams map[string]exchange.ExchangeService_ExchangeMsgServer
}

func (exchangeService *ExchangeService) ExchangeMsg(msgServer exchange.ExchangeService_ExchangeMsgServer) error {
	md, _ := metadata.FromIncomingContext(msgServer.Context())
	gatewayIP := md.Get("X-Gateway-IP")
	if len(gatewayIP) == 0 || gatewayIP[0] == "" {
		return errors.New("gateway ip header not found")
	}

	exchangeService.streams[gatewayIP[0]] = msgServer
	log.Infof("gateway %s gRPC streaming connect\n", gatewayIP)

	defer func() {
		log.Infof("gateway %s gRPC streaming closed\n", gatewayIP)
		delete(exchangeService.streams, gatewayIP[0])
	}()

	for {
		msg, err := msgServer.Recv()
		if errors.Is(err, io.EOF) {
			log.Infof("gateway stream is closed\n")
			break
		}
		if err != nil {
			log.Infof("recv err: %s\n", err.Error())
			break
		}
		if msg.Payload == nil {
			log.Infof("recv payload must not nil\n")
			continue
		}
		if err := dispatcher.GetMsgHandler(msg.Payload.Type).HandleProtocol(msg.Payload); err != nil {
			log.Errorf("handler msg failed. %+v", msg.Payload)
		}
	}

	return nil
}
