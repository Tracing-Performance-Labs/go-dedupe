package dedupe

type table[T Repr] interface {
	// This methods finds an appropriate representation for the string s and returns it.
	StoreAsRepr(s string) T

	// Recovers the original object from its representation.
	RecoverFromRepr(repr T) (string, error)
}
