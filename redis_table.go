package dedupe

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

// A table that uses Redis as a backend store.
type redisTable struct {
	client *redis.Client
	ctx    context.Context
}

// Create a new Redis backed table.
func NewRedisTable() *redisTable {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisAddr := redisHost + ":" + redisPort

	return &redisTable{
		client: redis.NewClient(&redis.Options{
			Addr: redisAddr,
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
