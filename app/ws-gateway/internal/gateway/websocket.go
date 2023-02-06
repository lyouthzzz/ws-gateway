package gateway

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/lyouthzzz/ws-gateway/api/wsapi"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/mapping"
	metrics "github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/metric"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/socketid"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/upstream"
	"github.com/lyouthzzz/ws-gateway/pkg/netutil"
	"net/http"
	"sync"
)

func NewWebsocketGateway(opts ...WebsocketGatewayOption) *WebsocketGateway {
	gateway := &WebsocketGateway{
		sidGenerator:    socketid.NewAtomicGenerator(0),
		upgrader:        &websocket.Upgrader{},
		sidConnRelation: mapping.NewSidConnMapping(),
		localIP:         netutil.LocalIPString(),
	}

	for _, opt := range opts {
		opt(gateway)
	}

	go gateway.recvMsg()

	return gateway
}

type WebsocketGateway struct {
	upstream        upstream.Upstream
	upgrader        *websocket.Upgrader
	sidGenerator    socketid.Generator
	sidConnRelation *mapping.SidConnMapping
	localIP         string

	mu sync.Mutex
}

func (gateway *WebsocketGateway) WebsocketConnectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := gateway.upgrader.Upgrade(w, r, nil)
		if err != nil {
			_, _ = w.Write([]byte("websocket upgrade err: " + err.Error()))
			return
		}

		metrics.GatewayOnlineTotals.Inc()

		nextSid, _ := gateway.sidGenerator.NextSid()
		gateway.sidConnRelation.Add(nextSid, conn)

		clear := func() {
			metrics.GatewayOnlineTotals.Dec()
			gateway.sidConnRelation.Delete(nextSid)
			_ = conn.Close()

			log.Infof("conn closed. sid: %d\n", nextSid)
		}
		defer clear()

		log.Infof("conn connect. sid: %d", nextSid)

		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				log.Infof("read message err: %s", err.Error())
				break
			}

			metrics.GatewayInputBytes.Add(float64(len(data)))

			protoMsg := &wsgateway.Protocol{}
			if err := jsoniter.Unmarshal(data, &protoMsg); err != nil {
				log.Infof("unmarshal protocol err: %s", err.Error())
				continue
			}

			if err = gateway.upstream.Send(&wsapi.Msg{
				Sid:       nextSid,
				GatewayIP: gateway.localIP,
				Payload:   protoMsg,
			}); err != nil {
				log.Infof("send upstream err: %s", err.Error())
			}
		}
	}
}

func (gateway *WebsocketGateway) recvMsg() {
	for {
		msg, err := gateway.upstream.Recv()
		if err != nil {
			log.Infof("recv msg err: %s", err.Error())
			continue
		}
		if msg.Sid == 0 {
			log.Infof("recv msg sid is 0, continue")
			continue
		}

		conn, has := gateway.sidConnRelation.Get(msg.Sid)
		if !has {
			log.Infof("conn offline. sid: %d", msg.Sid)
			continue
		}
		data, _ := jsoniter.Marshal(msg.Payload)
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Infof("conn write err: %s", err.Error())
			continue
		}

		metrics.GatewayOutputBytes.Add(float64(len(data)))
	}
}
