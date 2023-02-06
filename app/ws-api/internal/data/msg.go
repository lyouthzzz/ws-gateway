package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/domain"
)

var _ domain.MsgRepo = (*msgRepo)(nil)

type msgRepo struct {
	rdb *redis.Client
}

func (m *msgRepo) PublishMsg(ctx context.Context, msg *domain.Msg) error {
	return m.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: msg.Type,
		Values: msg.Payload,
	}).Err()
}
