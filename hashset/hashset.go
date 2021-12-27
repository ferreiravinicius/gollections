package hashset

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

func (set HashSet[T]) Add(item T) bool {
	beforeLen := len(set)
	set[item] = struct{}{}
	return beforeLen < len(set)
}

func (set HashSet[T]) Len() int {
	return len(set)
}

func (set HashSet[T]) AddAll(items ...T) bool {
	beforeLen := len(set)
	for _, el := range items {
		set[el] = struct{}{}
	}
	return beforeLen < len(set)
}

func (set HashSet[T]) Remove(item T) bool {
	if _, exists := set[item]; exists {
		delete(set, item)
		return true
	}
	return false
}

func (set HashSet[T]) RemoveAll(items ...T) bool {
	beforeLen := len(set)
	for _, el := range items {
		delete(set, el)
	}
	return beforeLen > len(set)
}

type Consumer[T any] func(item T) 

func (set HashSet[T]) ForEach(consumer Consumer[T]) {
	for item := range set {
		consumer(item)
	}
}
