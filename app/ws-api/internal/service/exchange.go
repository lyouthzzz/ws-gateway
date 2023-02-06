package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lyouthzzz/ws-gateway/api/wsapi"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/biz"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/domain"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"io"
)

const _defaultRecvChanCap = 10000

var ProviderSet = wire.NewSet(NewExchangeService)

func NewExchangeService() *ExchangeService {
	svc := &ExchangeService{streams: make(map[string]wsapi.ExchangeService_ExchangeMsgServer), recvMsgChan: make(chan *wsapi.Msg, _defaultRecvChanCap)}
	go svc.sendMsg()
	return svc
}

type ExchangeService struct {
	wsapi.UnimplementedExchangeServiceServer
	streams     map[string]wsapi.ExchangeService_ExchangeMsgServer
	exchangeBiz *biz.ExchangeBiz
	recvMsgChan chan *wsapi.Msg
}

func (svc *ExchangeService) ExchangeMsg(msgServer wsapi.ExchangeService_ExchangeMsgServer) error {
	md, _ := metadata.FromIncomingContext(msgServer.Context())
	gatewayIP := md.Get("X-Gateway-IP")
	if len(gatewayIP) == 0 || gatewayIP[0] == "" {
		return errors.New("gateway ip header not found")
	}

	svc.streams[gatewayIP[0]] = msgServer
	log.Infof("gateway %s gRPC streaming connect\n", gatewayIP)

	defer func() {
		log.Infof("gateway %s gRPC streaming closed\n", gatewayIP)
		delete(svc.streams, gatewayIP[0])
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

		svc.recvMsgChan <- msg
	}

	return nil
}

func (svc *ExchangeService) sendMsg() {
	for msg := range svc.recvMsgChan {
		if err := svc.exchangeBiz.Send(context.Background(),
			&domain.Msg{Type: msg.Payload.Type, Payload: msg.Payload.Payload},
		); err != nil {
			log.Errorf("exchange biz send error %s", err.Error())
		}
	}
}

func (svc *ExchangeService) Close() {
	close(svc.recvMsgChan)
}
