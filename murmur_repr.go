package dedupe

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/spaolacci/murmur3"
)

type murmurRepr struct{}

func NewMurmurRepr() ObjectRepr[string] {
	return &murmurRepr{}
}

func (r *murmurRepr) GetRepr(s string) string {
	bytes := []byte(s)
	hash := murmur3.Sum32(bytes)

	output := make([]byte, 4)
	binary.BigEndian.PutUint32(output, hash)

	repr := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(output)

	if len(repr) > len(s) {
		return s
	}

	return repr
}
