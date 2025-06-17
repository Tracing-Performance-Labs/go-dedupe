package dedupe

// Configure the Codec to use redis as the backing table.
func WithRedisTable() CodecOption {
	return func(c *Codec) {
		c.t = NewRedisTable()
	}
}

// Configure the Codec to use the default object representation.
func WithDefaultObjectRepr() CodecOption {
	return func(c *Codec) {
		c.repr = NewDefaultObjectRepr()
	}
}

// Configure the Codec to use Murmur3 as the object representation.
func WithMurmurRepr() CodecOption {
	return func(c *Codec) {
		c.repr = &murmurRepr{}
	}
}
