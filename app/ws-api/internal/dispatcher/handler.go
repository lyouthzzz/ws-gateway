package dispatcher

import (
	"fmt"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway"
)

var msgHandlerRegistry map[string]MsgHandler

type MsgHandler interface {
	HandleProtocol(*wsgateway.Protocol) error
}

func GetMsgHandler(name string) MsgHandler {
	handler, ok := msgHandlerRegistry[name]
	if ok {
		return handler
	}
	return msgHandlerRegistry["DEFAULT"]
}

func RegisterHandler(name string, handler MsgHandler) {
	msgHandlerRegistry[name] = handler
}

func init() {
	msgHandlerRegistry = make(map[string]MsgHandler)
	msgHandlerRegistry["DEFAULT"] = &defaultMsgHandler{}
}

var _ MsgHandler = (*defaultMsgHandler)(nil)

type defaultMsgHandler struct{}

func (handler *defaultMsgHandler) HandleProtocol(p *wsgateway.Protocol) error {
	fmt.Println("unknown protocol type " + p.Type)
	return nil
}
