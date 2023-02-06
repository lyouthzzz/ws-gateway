package gateway

import (
	"github.com/go-kratos/kratos/v2/encoding"
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
		sidGenerator:   socketid.NewAtomicGenerator(0),
		upgrader:       &websocket.Upgrader{},
		sidConnMapping: mapping.NewSidConnMapping(),
		ip:             netutil.LocalIPString(),
	}

	for _, opt := range opts {
		opt(gateway)
	}

	go gateway.recvMsg()

	return gateway
}

type WebsocketGateway struct {
	upstream       upstream.Upstream
	wsAPIClient    wsapi.ExchangeServiceClient
	upgrader       *websocket.Upgrader
	sidGenerator   socketid.Generator
	sidConnMapping *mapping.SidConnMapping
	ip             string

	mu sync.Mutex
}

func (gateway *WebsocketGateway) WebsocketConnectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := gateway.upgrader.Upgrade(w, r, nil)
		if err != nil {
			// todo
			_, _ = w.Write([]byte("websocket upgrade err: " + err.Error()))
			return
		}

		token := r.Header.Get("token")
		if token == "" {
			// todo
			_, _ = w.Write([]byte("websocket no authorization"))
			return
		}

		metrics.GatewayOnlineTotals.Inc()

		sid, _ := gateway.sidGenerator.NextSid()
		gateway.sidConnMapping.Add(sid, conn)

		// 用户连接
		connectReply, err := gateway.wsAPIClient.Connect(r.Context(), &wsapi.ConnectRequest{
			Server: gateway.ip,
			Sid:    sid,
			Token:  []byte(token),
		})
		if err != nil {
			// todo
			return
		}
		log.Infof("user %d connected", connectReply.Uid)

		clear := func() {
			metrics.GatewayOnlineTotals.Dec()
			gateway.sidConnMapping.Delete(sid)
			// 用户断开连接
			_, _ = gateway.wsAPIClient.Disconnect(r.Context(), &wsapi.DisconnectRequest{Server: gateway.ip, Sid: sid})

			_ = conn.Close()

			log.Infof("conn closed. sid: %d\n", sid)
		}
		defer clear()

		log.Infof("conn connect. sid: %d", sid)
		// json 序列化
		codec := encoding.GetCodec("json")

		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				log.Errorf("read message err: %s", err.Error())
				break
			}

			metrics.GatewayInputBytes.Add(float64(len(data)))

			protoMsg := &wsgateway.Protocol{}
			if err := codec.Unmarshal(data, &protoMsg); err != nil {
				log.Errorf("unmarshal protocol err: %s", err.Error())
				continue
			}

			if err = gateway.upstream.Send(&wsapi.Msg{
				Server:  gateway.ip,
				Sid:     sid,
				Payload: protoMsg,
			}); err != nil {
				log.Errorf("send upstream err: %s", err.Error())
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
		if msg.Sid == "" {
			log.Infof("recv msg sid is empty, continue")
			continue
		}

		conn, has := gateway.sidConnMapping.Get(msg.Sid)
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
