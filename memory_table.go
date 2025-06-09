package dedupe

import "errors"

type MemoryTable[T Repr] struct {
	data map[string]T
}

func NewMemoryTable[T Repr]() *MemoryTable[T] {
	return &MemoryTable[T]{
		data: make(map[string]T),
	}
}

func (mt *MemoryTable[T]) Store(s string, repr T) {
	mt.data[s] = repr
}

func (mt *MemoryTable[T]) Lookup(s string) (T, error) {
	val, ok := mt.data[s]
	if !ok {
		var zero T
		return zero, errors.New("value not found in memory table")
	}
	return val, nil
}
