package dedupe

import (
	"testing"
)

func TestConformsToInterface(t *testing.T) {
	withRedis(t, func() {
		var table table[string]
		table = NewRedisTable()
		_ = table
	})
}

func TestStoringValueReturnsIt(t *testing.T) {
	withRedis(t, func() {
		tbl := NewRedisTable()

		for _, s := range testStrings {
			tbl.Store(s, s)
		}

		for _, s := range testStrings {
			if ss, err := tbl.Lookup(s); ss != s || err != nil {
				t.Errorf("Expected %s, got %s with error %v", s, ss, err)
			}
		}
	})
}
