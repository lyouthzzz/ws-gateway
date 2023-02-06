package upstream

import (
	"github.com/lyouthzzz/ws-gateway/api/wsapi"
)

// Upstream handle bi streaming reconnect by self
type Upstream interface {
	Recv() (*wsapi.Msg, error)
	Send(*wsapi.Msg) error
	Close() error
}
