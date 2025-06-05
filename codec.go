package dedupe

type Repr interface {
	int32 | float32
}

type Codec[T Repr] struct{}

func (c *Codec[T]) Encode(s string) T {
	return 0
}

func (c *Codec[T]) Decode(o T) string {
	return ""
}
