package hashset

import "github.com/ferreiravinicius/gollections/set"

type HashSet[T comparable] map[T]struct{}

func New[T comparable]() HashSet[T] {
	return make(map[T]struct{})
}

func WithCapacity[T comparable](initialCapacity int) HashSet[T] {
	return make(map[T]struct{}, initialCapacity)
}

func From[T comparable](items ...T) HashSet[T] {
	set := WithCapacity[T](len(items))
	set.AddAll(items...)
	return set
}

func (s HashSet[T]) Add(item T) bool {
	beforeLen := len(s)
	s[item] = struct{}{}
	return beforeLen < len(s)
}

func (s HashSet[T]) Len() int {
	return len(s)
}

func (s HashSet[T]) AddAll(items ...T) bool {
	beforeLen := len(s)
	for _, el := range items {
		s[el] = struct{}{}
	}
	return beforeLen < len(s)
}

func (s HashSet[T]) Remove(item T) bool {
	if _, exists := s[item]; exists {
		delete(s, item)
		return true
	}
	return false
}

func (s HashSet[T]) RemoveAll(items ...T) bool {
	beforeLen := len(s)
	for _, el := range items {
		delete(s, el)
	}
	return beforeLen > len(s)
}

func (s HashSet[T]) ForEach(executeAction set.Action[T]) {
	for item := range s {
		executeAction(item)
	}
}
