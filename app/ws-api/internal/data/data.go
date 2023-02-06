package data

import (
	"github.com/go-redis/redis/v8"
	appconfig "github.com/lyouthzzz/ws-gateway/app/ws-api/internal/config"
)

func NewRedisClient(c *appconfig.Data_Redis) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: c.Addr, Password: c.Password, DB: int(c.Db)})
}
