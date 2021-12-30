package collection

type Collection[T any] interface {
	// Returns the lenght of this collection.
	Len() int

	// Returns true if this collection has no items.
	IsEmpty() bool

	// Creates a new slice containing all items from this collection.
	ToSlice() []T

	// Returns true if provided item exists in this set.
	Contains(item T) bool
}
