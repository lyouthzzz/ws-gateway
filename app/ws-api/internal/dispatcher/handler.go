package dispatcher

import (
	"fmt"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway"
)

var msgHandlerRegistry map[string]MsgHandler

type MsgHandler interface {
	HandleProtocol(*wsgateway.Protocol) error
}

func GetMsgHandler(t string) MsgHandler {
	handler, ok := msgHandlerRegistry[t]
	if ok {
		return handler
	}
	return msgHandlerRegistry["default"]
}

func init() {
	msgHandlerRegistry = make(map[string]MsgHandler)
	msgHandlerRegistry["default"] = &defaultMsgHandler{}
}

var _ MsgHandler = (*defaultMsgHandler)(nil)

type defaultMsgHandler struct{}

func (handler *defaultMsgHandler) HandleProtocol(p *wsgateway.Protocol) error {
	fmt.Println("unknown protocol type " + p.Type)
	return nil
}
