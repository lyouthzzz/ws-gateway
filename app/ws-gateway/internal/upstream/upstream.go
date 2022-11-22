package upstream

import (
	"github.com/lyouthzzz/ws-gateway/app/ws-api/api/exchange"
)

// Upstream handle bi streaming reconnect by self
type Upstream interface {
	Recv() (*exchange.Msg, error)
	Send(*exchange.Msg) error
	Close() error
}
