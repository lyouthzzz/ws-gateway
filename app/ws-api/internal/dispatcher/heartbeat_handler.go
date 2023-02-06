package dispatcher

import "github.com/lyouthzzz/ws-gateway/api/wsgateway"

var _ MsgHandler = (*HeartBeatHandler)(nil)

type HeartBeatHandler struct{}

func (handler *HeartBeatHandler) HandleProtocol(protocol *wsgateway.Protocol) error {
	//TODO implement me
	panic("implement me")
}

func (handler *HeartBeatHandler) Name() string {
	//TODO implement me
	panic("implement me")
}
