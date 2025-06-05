package dedupe

type Repr interface {
}

// Defines a single method that allows clients of the interface to get a representation of the provided string.
type ObjectRepr[T Repr] interface {
	GetRepr(s string) T
	ParseRepr(s string) (T, error)
}

type simpleObjectRepr[T Repr] struct{}

func (r *simpleObjectRepr[T]) GetRepr(s string) T {
	// TODO
	var repr T
	return repr
}

func (r *simpleObjectRepr[T]) ParseRepr(s string) (T, error) {
	// TODO
	var repr T
	return repr, nil
}

// NewDefaultObjectRepr creates a new default implementation of ObjectRepr.
func NewDefaultObjectRepr[T Repr]() ObjectRepr[T] {
	return &simpleObjectRepr[T]{}
}
