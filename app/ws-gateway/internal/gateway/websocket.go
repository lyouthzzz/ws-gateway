package gateway

import (
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway/protocol"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/relation"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/socketid"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/upstream"
	"github.com/lyouthzzz/ws-gateway/pkg/netutil"
	"log"
	"net/http"
	"sync"
)

func NewWebsocketGateway(opts ...WebsocketGatewayOption) *WebsocketGateway {
	gateway := &WebsocketGateway{
		sidGenerator:    socketid.NewAtomicGenerator(0),
		upgrader:        &websocket.Upgrader{},
		sidConnRelation: relation.NewSidConnRelation(),
		localIP:         netutil.LocalIPString(),

		logger:          log.Default(),
	}

	for _, opt := range opts {
		opt(gateway)
	}

	return gateway
}

type WebsocketGateway struct {
	upstream        upstream.Upstream
	upgrader        *websocket.Upgrader
	sidGenerator    socketid.Generator
	sidConnRelation *relation.SidConnRelation
	localIP         string

	mu     sync.Mutex
	logger *log.Logger
}

func (gateway *WebsocketGateway) WebsocketConnectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := gateway.upgrader.Upgrade(w, r, nil)
		if err != nil {
			_, _ = w.Write([]byte("websocket upgrade err: " + err.Error()))
			return
		}

		nextSid, _ := gateway.sidGenerator.NextSid()
		gateway.sidConnRelation.Add(nextSid, conn)

		clear := func() {
			gateway.sidConnRelation.Delete(nextSid)
			_ = conn.Close()
			gateway.logger.Printf("conn closed. sid: %d\n", nextSid)
		}
		defer clear()

		gateway.logger.Printf("conn connect. sid: %d\n", nextSid)

		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				gateway.logger.Printf("read message err: %s\n", err.Error())
				break
			}

			protoMsg := &protocol.Protocol{}
			if err := jsoniter.Unmarshal(data, &protoMsg); err != nil {
				gateway.logger.Printf("unmarshal protocol err: %s\n", err.Error())
				continue
			}

			if err = gateway.upstream.Send(&exchange.Msg{
				Sid:       nextSid,
				GatewayIP: gateway.localIP,
				Payload:   nil,
			}); err != nil {
				gateway.logger.Printf("send upstream err: %s\n", err.Error())
			}
		}
	}
}

func (gateway *WebsocketGateway) recvMsg() {
	for {
		msg, err := gateway.upstream.Recv()
		if err != nil {
			gateway.logger.Printf("recv msg err: %s\n", err.Error())
			continue
		}
		if msg.Sid == 0 {
			gateway.logger.Printf("recv msg sid is 0, continue")
			continue
		}

		conn, has := gateway.sidConnRelation.Get(msg.Sid)
		if !has {
			gateway.logger.Printf("conn offline. sid: %d\n", msg.Sid)
			continue
		}
		data, _ := jsoniter.Marshal(msg.Payload)
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			gateway.logger.Printf("conn write err: %s", err.Error())
		}
	}
}
