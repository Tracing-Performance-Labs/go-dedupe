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

var reprs = []ObjectRepr[string]{
	NewDefaultObjectRepr(),
	&murmurRepr{},
}

func TestGetRepr(t *testing.T) {
	for _, r := range reprs {
		for _, s := range testStrings {
			repr := r.GetRepr(s)
			if len(repr) > len(s) {
				t.Errorf("Expected the size of the represenation[%s] to be less than or equal to the size of the original value[%s]",
					repr, s)
			}
		}
	}
}

var sink string

func BenchmarkDefaultRepr(b *testing.B) {
	r := NewDefaultObjectRepr()
	b.RunParallel(func(b *testing.PB) {
		for b.Next() {
			for _, s := range testStrings {
				sink = r.GetRepr(s)
			}
		}
	})
}
