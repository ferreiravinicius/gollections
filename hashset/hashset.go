package hashset

type HashSet[T comparable] struct {
	hashMap map[T]struct{}
}

func New[T comparable](size ...int) *HashSet[T] {
	if len(size) == 1 {
		return &HashSet[T]{ hashMap: make(map[T]struct{}, size[0]) }
	}
	return &HashSet[T]{ hashMap: make(map[T]struct{}) }
}

func From[T comparable](items ...T) *HashSet[T] {
	set := New[T](len(items))
	set.AddAll(items...)
	return set
}

func (set HashSet[T]) Add(item T) bool {
	beforeLen := len(set.hashMap)
	set.hashMap[item] = struct{}{}
	return beforeLen < len(set.hashMap)
}

func (set HashSet[T]) Len() int {
	return len(set.hashMap)
}

func (set HashSet[T]) AddAll(items ...T) bool {
	beforeLen := len(set.hashMap)
	for _, el := range items {
		set.hashMap[el] = struct{}{}
	}
	return beforeLen < len(set.hashMap)
}

func (set HashSet[T]) Remove(item T) bool {
	if _, exists := set.hashMap[item]; exists {
		delete(set.hashMap, item)
		return true
	}
	return false
}

func (set HashSet[T]) RemoveAll(items ...T) bool {
	beforeLen := len(set.hashMap)
	for _, el := range items {
		delete(set.hashMap, el)
	}
	return beforeLen > len(set.hashMap)
}
