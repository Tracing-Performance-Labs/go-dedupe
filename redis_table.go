package dedupe

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// A table that uses Redis as a backend store.
type redisTable[T Repr] struct {
	repr   ObjectRepr[T]
	client *redis.Client
	ctx    context.Context
}

// Create a new Redis backed table.
func NewRedisTable[T Repr](repr ObjectRepr[T]) *redisTable[T] {
	return &redisTable[T]{
		repr: repr,
		client: redis.NewClient(&redis.Options{
			// TODO: Get from enviroment variable.
			// TODO: Look into how go services are normally configured.
			Addr: "localhost:6379",
		}),
		ctx: context.Background(),
	}
}

func (rt *redisTable[T]) Lookup(s string) T {
	val, err := rt.client.Get(rt.ctx, s).Result()
	if err != nil {
		// TODO: Compute representation.
		// TODO: Store representation in Redis.
		// TODO: Set val.
	}

	repr, err := rt.repr.ParseRepr(val)
	if err != nil {
		panic("Failed to parse representation: " + err.Error())
	}

	return repr
}
