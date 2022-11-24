package upstream

import (
	"context"
	"github.com/lyouthzzz/ws-gateway/api/wsapi/exchange"
	"github.com/lyouthzzz/ws-gateway/pkg/netutil"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
)

var (
	_gRPCStreamingUpstreamDefaultRecvMsgChanCap = 10000
	_gRPCStreamingUpstreamDefaultSendMsgChanCap = 10000
)

var _ Upstream = (*gRPCStreamingUpstream)(nil)

type GRPCStreamingUpstreamOption func(*gRPCStreamingUpstream)

func GRPCStreamingExchangeClient(svr exchange.ExchangeServiceClient) GRPCStreamingUpstreamOption {
	return func(upstream *gRPCStreamingUpstream) { upstream.exc = svr }
}

func GRPCStreamingUpstreamOptionRecvMsgChanCap(cap int) GRPCStreamingUpstreamOption {
	return func(upstream *gRPCStreamingUpstream) { upstream.recvMsgChanCap = cap }
}

func GRPCStreamingUpstreamOptionSendMsgChanCap(cap int) GRPCStreamingUpstreamOption {
	return func(upstream *gRPCStreamingUpstream) { upstream.sendMsgChanCap = cap }
}

func NewGRPCStreamingUpstream(opts ...GRPCStreamingUpstreamOption) (Upstream, error) {
	up := &gRPCStreamingUpstream{
		recvMsgChanCap: _gRPCStreamingUpstreamDefaultRecvMsgChanCap,
		sendMsgChanCap: _gRPCStreamingUpstreamDefaultSendMsgChanCap,
		logger:         log.Default(),
	}

	for _, opt := range opts {
		opt(up)
	}

	var err error
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"X-Gateway-IP": netutil.LocalIPString()}))
	if up.msgc, err = up.exc.ExchangeMsg(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	up.recvMsgChan = make(chan *exchange.Msg, up.recvMsgChanCap)
	up.sendMsgChan = make(chan *exchange.Msg, up.sendMsgChanCap)

	go up.sendMsg()
	go up.recvMsg()

	return up, nil
}

type gRPCStreamingUpstream struct {
	exc  exchange.ExchangeServiceClient
	msgc exchange.ExchangeService_ExchangeMsgClient

	recvMsgChanCap int
	sendMsgChanCap int
	recvMsgChan    chan *exchange.Msg
	sendMsgChan    chan *exchange.Msg

	// todo ??? 0: init 1: running 2: reconnection 3: stopping 4: stopped
	status int32
	// todo ??? 0: reconnecting 1: closed
	events chan int

	logger *log.Logger
}

func (upstream *gRPCStreamingUpstream) Recv() (*exchange.Msg, error) {
	return <-upstream.recvMsgChan, nil
}

func (upstream *gRPCStreamingUpstream) Send(msg *exchange.Msg) error {
	upstream.sendMsgChan <- msg
	return nil
}

func (upstream *gRPCStreamingUpstream) Close() error {
	return upstream.msgc.CloseSend()
}

// todo ??? use chan buffer or send gRPC Streaming sync
func (upstream *gRPCStreamingUpstream) sendMsg() {
	for msg := range upstream.sendMsgChan {
		_ = upstream.msgc.SendMsg(msg)
	}
}

func (upstream *gRPCStreamingUpstream) recvMsg() {
	for {
		// todo handle
		//if atomic.LoadInt32(&upstream.status) != 1 {
		//	continue
		//}

		msg, err := upstream.msgc.Recv()
		if errors.Is(err, io.EOF) {
			// reconnect by yourself
			break
		}
		if err != nil {
			upstream.logger.Printf("gRPC streaming recv err: %s\n", err.Error())
			break
		}
		upstream.recvMsgChan <- msg
	}
}
