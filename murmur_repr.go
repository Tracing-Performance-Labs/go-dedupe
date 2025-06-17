package dedupe

import (
	"strconv"

	"github.com/spaolacci/murmur3"
)

type murmurRepr struct{}

func NewMurmurRepr() ObjectRepr[string] {
	return &murmurRepr{}
}

func (r *murmurRepr) GetRepr(s string) string {
	bytes := []byte(s)
	hash := murmur3.Sum32(bytes)
	repr := strconv.FormatUint(uint64(hash), 10)

	if len(repr) > len(s) {
		return s
	}

	return repr
}
