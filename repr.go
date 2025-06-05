package dedupe

type Repr interface {
	int32 | float32
}

// Defines a single method that allows clients of the interface to get a representation of the provided string.
type ObjectRepr[T Repr] interface {
	GetRepr(s string) T
}

type simpleObjectRepr[T Repr] struct{}

func (r *simpleObjectRepr[T]) GetRepr(s string) T {
	// TODO
	return 0
}

// NewDefaultObjectRepr creates a new default implementation of ObjectRepr.
func NewDefaultObjectRepr[T Repr]() ObjectRepr[T] {
	return &simpleObjectRepr[T]{}
}
