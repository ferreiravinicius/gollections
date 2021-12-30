package collection

import "fmt"

type Collection[T any] interface {
	fmt.Stringer
	// Returns the lenght of this collection.
	Len() int

	// Returns true if this collection has no items.
	IsEmpty() bool

	// Creates a new slice containing all items from this collection.
	ToSlice() []T
}
