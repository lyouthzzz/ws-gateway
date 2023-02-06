package biz

import (
	"context"
	"github.com/lyouthzzz/ws-gateway/api/wsgateway"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/domain"
)

type ExchangeBiz struct {
	msgRepo domain.MsgRepo
}

func (biz *ExchangeBiz) Send(ctx context.Context, msg *domain.Msg) error {
	if msg.Type == wsgateway.Type_name[wsgateway.Type_HEARTBEAT] {

	}
	return biz.msgRepo.PublishMsg(ctx, msg)
}
