package dedupe

import "testing"

var testStrings = []string{
	"hello",
	"world",
	"black",
	"superfluous",
	"yes",
	"marvelous",
	"there",
	"the",
	"Firefox",
	"http.port.dest",
	"no",
	"a",
}

func TestGetRepr(t *testing.T) {
	simpleRepr := NewDefaultObjectRepr()

	for _, s := range testStrings {
		repr := simpleRepr.GetRepr(s)
		if len(repr) > len(s) {
			t.Errorf("Expected the size of the represenation to be less than or equal to the size of the original value")
		}
	}
}
