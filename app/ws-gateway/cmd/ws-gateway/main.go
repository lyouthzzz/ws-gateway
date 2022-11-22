package main

import (
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/gateway"
	"net/http"
)

func main() {
	websocketGateway := gateway.NewWebsocketGateway()

	http.HandleFunc("/gateway/ws", websocketGateway.WebsocketConnectHandler())
}
