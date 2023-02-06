package dispatcher

import (
	"fmt"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway"
)

var msgHandlerRegistry map[string]MsgHandler

type MsgHandler interface {
	HandleProtocol(*wsgateway.Protocol) error
	Name() string
}

func GetMsgHandler(name string) MsgHandler {
	handler, ok := msgHandlerRegistry[name]
	if ok {
		return handler
	}
	return globalDefaultMsgHandler
}

func RegisterHandler(handler MsgHandler) {
	msgHandlerRegistry[handler.Name()] = handler
}

func init() {
	msgHandlerRegistry = make(map[string]MsgHandler)
	msgHandlerRegistry["DEFAULT"] = &defaultMsgHandler{}
}

var (
	_                       MsgHandler = (*defaultMsgHandler)(nil)
	globalDefaultMsgHandler            = &defaultMsgHandler{}
)

type defaultMsgHandler struct{}

func (handler *defaultMsgHandler) Name() string {
	return "DEFAULT"
}

func (handler *defaultMsgHandler) HandleProtocol(p *wsgateway.Protocol) error {
	fmt.Println("unknown protocol type " + p.Type)
	return nil
}
