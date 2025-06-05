package dedupe

type Codec[T Repr] struct {
	t table[T]
}

func NewCodec[T Repr](t table[T]) *Codec[T] {
	return &Codec[T]{t: t}
}

func (c *Codec[T]) Encode(s string) T {
	return 0
}

func (c *Codec[T]) Decode(o T) string {
	return ""
}
