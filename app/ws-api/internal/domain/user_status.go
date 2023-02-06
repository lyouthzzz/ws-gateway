package domain

import (
	"context"
	"strconv"
	"time"
)

type UserStatusRepo interface {
	Connect(ctx context.Context, status *UserStatus) error
	Disconnect(ctx context.Context, status *UserStatus) error
	KeepAlive(ctx context.Context, status *UserStatus) error
}

type UserStatus struct {
	// 网关IP
	Server string
	// socketId
	Sid string
	// 用户ID
	Uid uint64
	// 过期时间
	ExpireDuration time.Duration
}

func (status *UserStatus) Key() string {
	return strconv.FormatUint(status.Uid, 10)
}

func (status *UserStatus) Value() string {
	return ""
}
