package dedupe

type Codec[T Repr] struct {
	t table[T]
}

func NewCodec[T Repr](t table[T]) *Codec[T] {
	// TODO: Use options pattern to configure the codec.
	return &Codec[T]{t: t}
}

// Encode the provided value by obtaining a compact representation for it.
func (c *Codec[T]) Encode(s string) T {
	return c.t.StoreAsRepr(s)
}

// Retrieve a value from its compact representation.
func (c *Codec[T]) Decode(o T) string {
	// TODO
	return ""
}
