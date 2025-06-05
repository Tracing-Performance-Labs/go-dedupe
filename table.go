package dedupe

type table[T Repr] interface {
	// This methods finds an appropriate representation for the string s and returns it.
	Lookup(s string) T
}
