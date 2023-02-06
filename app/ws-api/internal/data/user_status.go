package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/domain"
	"strconv"
)

var _ domain.UserStatusRepo = (*userStatusRepo)(nil)

type userStatusRepo struct {
	rdb *redis.Client
}

func (repo *userStatusRepo) Connect(ctx context.Context, status *domain.UserStatus) error {
	return repo.rdb.Set(ctx, "user:online:"+strconv.FormatUint(status.Uid, 10), status.Value(), status.ExpireDuration).Err()
}

func (repo *userStatusRepo) Disconnect(ctx context.Context, status *domain.UserStatus) error {
	return repo.rdb.Del(ctx, "user:online:"+strconv.FormatUint(status.Uid, 10)).Err()
}

func (repo *userStatusRepo) KeepAlive(ctx context.Context, status *domain.UserStatus) error {
	return repo.rdb.Expire(ctx, "user:online:"+strconv.FormatUint(status.Uid, 10), status.ExpireDuration).Err()
}
