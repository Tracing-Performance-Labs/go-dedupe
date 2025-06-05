package dedupe

// A table that uses Redis as a backend store.
type redisTable[T Repr] struct {
	repr ObjectRepr[T]
}

// Create a new Redis backed table.
func NewRedisTable[T Repr](repr ObjectRepr[T]) *redisTable[T] {
	return &redisTable[T]{repr: repr}
}

func (rt *redisTable[T]) Lookup(s string) T {
	return 0
}
