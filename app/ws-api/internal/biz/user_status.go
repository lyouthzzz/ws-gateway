package biz

import (
	"context"
	"github.com/lyouthzzz/ws-gateway/app/ws-api/internal/domain"
)

type UserStatusBiz struct {
	userStatusRepo domain.UserStatusRepo
}

func (biz *UserStatusBiz) Connect(ctx context.Context, status *domain.UserStatus) error {
	return biz.userStatusRepo.Connect(ctx, status)
}

func (biz *UserStatusBiz) Disconnect(ctx context.Context, status *domain.UserStatus) error {
	return biz.userStatusRepo.Disconnect(ctx, status)
}

func (biz *UserStatusBiz) KeepAlive(ctx context.Context, status *domain.UserStatus) error {
	return biz.userStatusRepo.KeepAlive(ctx, status)
}
