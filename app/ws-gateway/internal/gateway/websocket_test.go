package gateway

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway/protocol"
	"testing"
)

func TestWebsocketProtocol(t *testing.T) {
	payload := map[string]string{
		"key": "value",
	}
	payloadContent, _ := jsoniter.Marshal(payload)

	protoMsg := &protocol.Protocol{
		Type:    "message",
		Payload: payloadContent,
	}

	protoMsgContent, _ := jsoniter.Marshal(protoMsg)
	fmt.Println(string(protoMsgContent))
}
