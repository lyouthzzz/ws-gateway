package domain

import "context"

type MsgRepo interface {
	PublishMsg(ctx context.Context, msg *Msg) error
}

type Msg struct {
	Type    string
	Payload []byte
}
