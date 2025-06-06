package dedupe

import (
	"testing"
)

var (
	repr = NewDefaultObjectRepr()
)

func TestConformsToInterface(t *testing.T) {
	var table table[string]
	table = NewRedisTable(repr)
	_ = table
}

func TestRedisStoreAsRepr(t *testing.T) {
	table := NewRedisTable(repr)

	for _, s := range testStrings {
		r := table.StoreAsRepr(s)
		if len(r) > len(s) {
			t.Errorf("Lookup returned a representation that is thiccer than the original string: %s | %s",
				s,
				r)
		}
	}
}

func TestRedisRecoverFromRepr(t *testing.T) {
	table := NewRedisTable(repr)

	for _, s := range testStrings {
		r := table.StoreAsRepr(s)
		o, err := table.RecoverFromRepr(r)
		if err != nil {
			t.Errorf("Failed to recover representation: %s", err.Error())
		}
		if o != s {
			t.Errorf("Recovered something other than the original string: %v %v",
				[]byte(s),
				[]byte(o))
		}
	}
}
