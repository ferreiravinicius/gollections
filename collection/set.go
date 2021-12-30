package collection

// A collection that contains no duplicate elements.
// This interface exists in case the user wants to
// depends on a contract instead of the implementation.
type Set[T any] interface {
	Collection[T]

	// Adds a new item to this set if it is not present.
	// Returns true if the item is added.
	Add(item T) bool

	// Adds all items to this set if it is not present.
	// Returns true if at least one item is added.
	AddAll(items ...T) bool

	// Removes from this set the specified item if it is present.
	// Returns true if item is removed.
	Remove(item T) bool

	// Removes from this set all items that are present.
	// Returns true if at least one item is removed.
	RemoveAll(items ...T) bool

	// Performs an action for each item inside this set.
	// A action is -> func(item T)
	ForEach(action Action[T])
}

// Represents the signature for a action.
// Same as -> func(item T)
type Action[T any] func(item T)
