package dedupe

type Codec[T Repr] struct {
	t table[T]
}

func NewCodec[T Repr](t table[T]) *Codec[T] {
	// TODO: Use options pattern to configure the codec.
	return &Codec[T]{t: t}
}

func (c *Codec[T]) Encode(s string) T {
	// TODO
	return 0
}

func (c *Codec[T]) Decode(o T) string {
	// TODO
	return ""
}
