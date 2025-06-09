package dedupe

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// A table that uses Redis as a backend store.
type redisTable struct {
	client *redis.Client
	ctx    context.Context
}

// Create a new Redis backed table.
func NewRedisTable() *redisTable {
	return &redisTable{
		client: redis.NewClient(&redis.Options{
			// TODO: Get from enviroment variable.
			// TODO: Look into how go services are normally configured.
			Addr: "localhost:6379",
		}),
		ctx: context.Background(),
	}
}

func (rt *redisTable) Lookup(s string) (string, error) {
	val, err := rt.client.Get(rt.ctx, s).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rt *redisTable) Store(s string, repr string) {
	status := rt.client.Set(rt.ctx, s, repr, 0)
	if status.Err() != nil {
		panic("Failed to set value in Redis: " + status.Err().Error())
	}
}
