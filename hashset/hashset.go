package hashset

import (
	"fmt"
	"strings"

	"github.com/ferreiravinicius/gollections/collection"
)

// HashSet is a collection that contains no duplicate elements.
// Implemention of set backed by a Hash Table (Go builtin Map).
// Order of insertion is not guaranteed.
// This implementation is not synchronized.
// It does implements the set.Set interface.
//
// Example:
//
// s := hashset.New()
type HashSet[T comparable] map[T]struct{}

// Creates a new default HashSet instance.
// Also see WithCapacity() and From.
//
// Example: hashset.New[int]()
func New[T comparable]() HashSet[T] {
	return make(map[T]struct{})
}

// Creates a new HashSet with specified capacity.
//
// Example: hashset.WithCapacity(100) //initial size of 100
func WithCapacity[T comparable](initialCapacity int) HashSet[T] {
	return make(map[T]struct{}, initialCapacity)
}

// Creates a new HashSet containg all provided items.
//
// Examples:
//
// s := hashset.From(1, 2, 3)
//
// items := []int{1, 2, 3}
//
// s := hashset.From(items...)
func From[T comparable](items ...T) HashSet[T] {
	set := WithCapacity[T](len(items))
	set.AddAll(items...)
	return set
}

// Adds a new item to this set if it is not present - O(1).
// Returns true if the item is added.
func (s HashSet[T]) Add(item T) bool {
	beforeLen := len(s)
	s[item] = struct{}{}
	return beforeLen < len(s)
}

// Adds all items to this set if it is not present.
// Returns true if at least one item is added.
func (s HashSet[T]) AddAll(items ...T) bool {
	beforeLen := len(s)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return beforeLen < len(s)
}

// Returns the size of this set.
func (s HashSet[T]) Len() int {
	return len(s)
}

// Removes from this set the specified item if it is present - O(1).
// Returns true if item is removed.
func (s HashSet[T]) Remove(item T) bool {
	if _, exists := s[item]; exists {
		delete(s, item)
		return true
	}
	return false
}

// Removes from this set all items that are present.
// Returns true if at least one item is removed.
func (s HashSet[T]) RemoveAll(items ...T) bool {
	beforeLen := len(s)
	for _, el := range items {
		delete(s, el)
	}
	return beforeLen > len(s)
}

// Performs an action for each item inside this set.
func (s HashSet[T]) ForEach(action collection.Action[T]) {
	for item := range s {
		action(item)
	}
}

func (s HashSet[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s HashSet[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s HashSet[T]) ToSlice() []T {
	result := make([]T, len(s))
	n := 0
	for item := range s {
		result[n] = item
		n++
	}
	return result
}

const (
	prefix   = "HashSet{"
	suffix   = " }"
	template = " %v"
)

func (s HashSet[T]) String() string {
	var sb strings.Builder
	sb.WriteString(prefix)
	for item := range s {
		fmt.Fprintf(&sb, template, item)
	}
	sb.WriteString(suffix)
	return sb.String()
}
