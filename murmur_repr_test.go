package dedupe

import "testing"

func BenchmarkMurmurRepr(b *testing.B) {
	r := NewMurmurRepr()
	b.RunParallel(func(b *testing.PB) {
		for b.Next() {
			for _, s := range testStrings {
				sink = r.GetRepr(s)
			}
		}
	})
}
