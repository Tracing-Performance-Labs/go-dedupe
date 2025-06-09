package dedupe

import (
	"testing"
)

func TestCodecEncode(t *testing.T) {
	repr := NewDefaultObjectRepr()
	table := NewRedisTable()
	codec := NewCodec(table, repr)

	for _, s := range testStrings {
		r := codec.Encode(s)
		if len(r) > len(s) {
			t.Errorf("Encoded representation is not compact: %s %s", []byte(s), []byte(r))
		}
	}
}

func TestCodecDecode(t *testing.T) {
	repr := NewDefaultObjectRepr()
	table := NewRedisTable()
	codec := NewCodec(table, repr)

	for _, s := range testStrings {
		r := codec.Encode(s)
		o, err := codec.Decode(r)
		if err != nil {
			t.Errorf("Failed to decode representation: %s", err.Error())
		}
		if o != s {
			t.Errorf("Decoded something other than the original string: %v %v",
				[]byte(s),
				[]byte(o))
		}
	}
}
