package dedupe

import (
	"fmt"
	"strconv"
	"sync"
)

type Repr interface {
}

// Defines a single method that allows clients of the interface to get a representation of the provided string.
type ObjectRepr[T Repr] interface {
	GetRepr(s string) T
}

type simpleObjectRepr struct {
	counter int
	mtx     sync.Mutex
}

func (r *simpleObjectRepr) GetRepr(s string) string {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	count := r.counter
	r.counter += 1

	repr := fmt.Sprintf("%s", strconv.Itoa(count))

	if len(repr) > len(s) {
		return s
	}

	return repr
}

// NewDefaultObjectRepr creates a new default implementation of ObjectRepr.
func NewDefaultObjectRepr() ObjectRepr[string] {
	return &simpleObjectRepr{}
}
