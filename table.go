package dedupe

type table[T Repr] interface {
	// Store a value in the table.
	Store(s string, repr T)
	// Try to retrieve a value from the table.
	Lookup(s string) (T, error)
}
