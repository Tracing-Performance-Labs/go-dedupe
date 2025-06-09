package dedupe

// A codec can be used to convert a string into a a more compact representation.
type Codec struct {
	t    table[string]
	repr ObjectRepr[string]
}

func NewCodec(t table[string], repr ObjectRepr[string]) *Codec {
	// TODO: Use options pattern to configure the codec.
	return &Codec{
		t:    t,
		repr: repr,
	}
}

// Encode the provided value by obtaining a compact representation for it.
func (c *Codec) Encode(s string) string {
	val, err := c.t.Lookup(s)
	if err != nil {
		// 1. Compute representation.
		repr := c.repr.GetRepr(s)

		// 2. Store representation in the table.
		c.t.Store(s, repr)

		// 3. Store the reverse mapping.
		c.t.Store(repr, s)

		// 3. Set val.
		val = repr
	}
	return val
}

// Retrieve a value from its compact representation.
func (c *Codec) Decode(repr string) (string, error) {
	val, err := c.t.Lookup(repr)
	return val, err
}
