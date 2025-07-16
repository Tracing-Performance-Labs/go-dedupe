package dedupe

import (
	"sync"

	cuckoo "github.com/seiflotfy/cuckoofilter"
)

// A codec can be used to convert a string into a a more compact representation.
type Codec struct {
	t      table[string]
	repr   ObjectRepr[string]
	filter cuckoo.Filter
	cache  sync.Map
}

type CodecOption func(*Codec)

func NewCodec(opts ...CodecOption) *Codec {
	var (
		defaultTable = NewMemoryTable[string]()
		defaultRepr  = NewDefaultObjectRepr()
	)

	c := &Codec{
		t:      defaultTable,
		repr:   defaultRepr,
		filter: *cuckoo.NewFilter(1000000),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Codec) store(s string) string {
	// 1. Compute representation.
	repr := c.repr.GetRepr(s)

	// 2. Store representation in the table.
	c.t.Store(s, repr)

	// 3. Store the reverse mapping.
	c.t.Store(repr, s)

	// 3. Set val.
	return repr
}

// Encode the provided value by obtaining a compact representation for it.
func (c *Codec) Encode(s string) string {
	bs := []byte(s)
	seen := c.filter.Lookup(bs)
	if !seen {
		if !c.filter.Insert(bs) {
			panic("failed to store in filter")
		}
		val := c.store(s)
		c.cache.Store(s, val)
		return val
	}

	val, ok := c.cache.Load(s)
	if ok {
		return val.(string)
	}

	val, err := c.t.Lookup(s)
	if err != nil {
		return c.store(s)
	}

	return val.(string)
}

// Retrieve a value from its compact representation.
func (c *Codec) Decode(repr string) (string, error) {
	val, err := c.t.Lookup(repr)
	return val, err
}
