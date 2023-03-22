package gateway

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/lyouthzzz/ws-gateway/api/wsapi"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway"
	"github.com/lyouthzzz/ws-gateway/app/ws-gateway/internal/upstream"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"testing"
)

func TestWebsocketProtocol(t *testing.T) {
	v := randStr(1024 * 5)

	payload := map[string]string{
		"key": v,
	}
	payloadContent, _ := jsoniter.Marshal(payload)

	fmt.Printf("payload len: %d\n", len(payloadContent))

	protoMsg := &wsgateway.Protocol{
		Type:    "message",
		Payload: payloadContent,
	}

	protoMsgContent, _ := jsoniter.Marshal(protoMsg)

	fmt.Printf("payload len: %d\n", len(protoMsgContent))
	fmt.Println(string(protoMsgContent))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getMockGateway(b testing.TB) *WebsocketGateway {
	conn, err := grpc.DialContext(context.Background(), "localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(errors.WithStack(err))
	}
	exc := wsapi.NewExchangeServiceClient(conn)

	up, err := upstream.NewGRPCStreamingUpstream(
		upstream.GRPCStreamingExchangeClient(exc),
	)
	if err != nil {
		panic(errors.WithStack(err))
	}

	websocketGateway := NewWebsocketGateway(
		WebsocketGatewayOptionUpstream(up),
	)

	return websocketGateway
}

func BenchmarkWebsocketGateway_Send_case1(b *testing.B) {
	websocketGateway := getMockGateway(b)

	msg := &wsapi.Msg{
		Sid:    "1",
		Server: "localhost",
		Payload: &wsgateway.Protocol{
			Type:    "HEARTBEAT",
			Payload: []byte(randStr(1024)),
		},
	}
	msgContent, _ := jsoniter.Marshal(msg)
	fmt.Println(len(msgContent))

	for i := 0; i < b.N; i++ {
		err := websocketGateway.upstream.Send(msg)
		require.NoError(b, err)
	}
}
