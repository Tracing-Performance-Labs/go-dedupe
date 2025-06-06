package dedupe

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// A table that uses Redis as a backend store.
type redisTable struct {
	repr   ObjectRepr[string]
	client *redis.Client
	ctx    context.Context
}

// Create a new Redis backed table.
func NewRedisTable[T Repr](repr ObjectRepr[string]) *redisTable {
	return &redisTable{
		repr: repr,
		client: redis.NewClient(&redis.Options{
			// TODO: Get from enviroment variable.
			// TODO: Look into how go services are normally configured.
			Addr: "localhost:6379",
		}),
		ctx: context.Background(),
	}
}

func (rt *redisTable) Lookup(s string) string {
	val, err := rt.client.Get(rt.ctx, s).Result()
	if err != nil {
		// 1. Compute representation.
		repr := rt.repr.GetRepr(s)

		// 2. Store representation in Redis.
		rt.client.Set(rt.ctx, s, repr, 0)

		// 3. Set val.
		val = repr
	}

	return val
}
