package set

type Consumer[T any] func(item T)

type Set[T comparable] interface {
	Add(item T) bool
	AddAll(items ...T) bool
	Len() int
	Remove(item T) bool
	RemoveAll(items ...T) bool
	ForEach(consumer Consumer[T])
}
