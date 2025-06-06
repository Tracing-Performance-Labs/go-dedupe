package dedupe

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// A table that uses Redis as a backend store.
type redisTable struct {
	repr   ObjectRepr[string]
	client *redis.Client
	ctx    context.Context
}

// Create a new Redis backed table.
func NewRedisTable(repr ObjectRepr[string]) *redisTable {
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

func (rt *redisTable) StoreAsRepr(s string) string {
	val, err := rt.client.Get(rt.ctx, s).Result()
	if err != nil {
		// 1. Compute representation.
		repr := rt.repr.GetRepr(s)

		// 2. Store representation in Redis.
		status := rt.client.Set(rt.ctx, s, repr, 0)
		if status.Err() != nil {
			panic("Failed to set value in Redis: " + status.Err().Error())
		}

		status = rt.client.Set(rt.ctx, repr, s, 0)
		if status.Err() != nil {
			panic("Failed to set value in Redis: " + status.Err().Error())
		}

		// 3. Set val.
		val = repr
	}

	return val
}

func (rt *redisTable) RecoverFromRepr(repr string) (string, error) {
	val, err := rt.client.Get(rt.ctx, repr).Result()
	if err != nil {
		return "", errors.New(fmt.Sprintf(
			"Could not find representation: %v: %s", []byte(repr), err.Error()))
	}

	return val, nil
}
