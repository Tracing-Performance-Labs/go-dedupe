package dedupe

import (
	"testing"
)

func TestConformsToInterface(t *testing.T) {
	var table table[string]
	table = NewRedisTable()
	_ = table
}

// TODO: Test store

// TODO: Test lookup
