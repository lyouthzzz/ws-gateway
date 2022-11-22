package upstream

import (
	"github.com/lyouthzzz/ws-gateway/app/ws-api/api/exchange"
)

var _ Upstream = (*WebsocketUpstream)(nil)

type WebsocketUpstream struct {
}

func (upstream *WebsocketUpstream) Recv() (*exchange.Msg, error) {
	panic("implement me")
}

func (upstream *WebsocketUpstream) Send(msg *exchange.Msg) error {
	panic("implement me")
}

func (upstream *WebsocketUpstream) Close() error {
	panic("implement me")
}
