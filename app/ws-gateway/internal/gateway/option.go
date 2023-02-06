package gateway

import (
	"github.com/gorilla/websocket"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/socketid"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/upstream"
)

type WebsocketGatewayOption func(*WebsocketGateway)

func WebsocketGatewayOptionUpstream(up upstream.Upstream) WebsocketGatewayOption {
	return func(gateway *WebsocketGateway) { gateway.upstream = up }
}

func WebsocketGatewayOptionUpgrader(upgrader *websocket.Upgrader) WebsocketGatewayOption {
	return func(gateway *WebsocketGateway) { gateway.upgrader = upgrader }
}

func WebsocketGatewayOptionSidGenerator(g socketid.Generator) WebsocketGatewayOption {
	return func(gateway *WebsocketGateway) { gateway.sidGenerator = g }
}

func WebsocketGatewayOptionLocalIP(ip string) WebsocketGatewayOption {
	return func(gateway *WebsocketGateway) { gateway.ip = ip }
}
