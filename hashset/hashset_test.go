package hashset

import (
	"testing"
)

func eq[T comparable](t *testing.T, received T, expected T) {
	if expected != received {
		message := "expected: %v but got: %v"
		t.Errorf(message, expected, received)
	}
}

func contains[T comparable](slice []T, item T) bool {
	for _, each := range slice {
		if each == item {
			return true
		}	
	}
	return false
}

func TestHashSetNew(t *testing.T) {
	n := New[int]()
	eq(t, n.Len(), 0)
	withCap := WithCapacity[int](5)
	eq(t, withCap.Len(), 0)
}

func TestHashSetAdd(t *testing.T) {
	s := New[int]()
	eq(t, s.Add(1), true)
	eq(t, s.Len(), 1)
	eq(t, s.Add(1), false)
	eq(t, s.Len(), 1)
}

func TestHashSetLen(t *testing.T) {
	s := New[int]()
	eq(t, s.Len(), 0)
	s.Add(1)
	eq(t, s.Len(), 1)
	s.Add(1)
	eq(t, s.Len(), 1)
}

func TestHashSetAddAll(t *testing.T) {
	s := New[int]()

	eq(t, s.AddAll(), false)
	eq(t, s.Len(), 0)

	eq(t, s.AddAll(1, 2, 3), true)
	eq(t, s.Len(), 3)

	eq(t, s.AddAll(1, 2, 3, 4), true)
	eq(t, s.Len(), 4)

	eq(t, s.AddAll(1, 2), false)
	eq(t, s.Len(), 4)
}

func TestHashSetRemove(t *testing.T) {
	s := New[int]()
	eq(t, s.Remove(1), false)
	s.Add(1)
	eq(t, s.Remove(1), true)
	eq(t, s.Len(), 0)
}

func TestFrom(t *testing.T) {
	s1 := From[int](1, 2, 3, 4, 5)
	eq(t, s1.Len(), 5)

	slice := []int{1, 2, 3}
	s2 := From[int](slice...)
	eq(t, s2.Len(), 3)
}

func TestHashSetRemoveAll(t *testing.T) {
	s := From[int](1, 2, 3, 4, 5)

	eq(t, s.RemoveAll(), false)
	eq(t, s.Len(), 5)

	eq(t, s.RemoveAll(1, 2, 3), true)
	eq(t, s.Len(), 2)

	eq(t, s.RemoveAll(5, 6), true)
	eq(t, s.Len(), 1)
}



func (set HashSet[T]) IterItems() (chan T, func()) {

	results := make(chan T, len(set))
	cancel := make(chan bool, 1)
	exit := func() {
		cancel <- true
	}
	
	go func(results chan T, cancel chan bool) {
		defer close(results)
		defer close(cancel)
		for item := range set {
			curr := item
			select {
			case results <- curr:
			case <- cancel:
				break
			}	
		}
	}(results, cancel)

	return results, exit
}

func TestHashSetForEach(t *testing.T) {
	set := From[int](1, 2, 3)
	r := make([]int, 0, 3)
	set.ForEach(func(item int) {
		r = append(r, item)
	})
	eq(t, contains(r, 1), true)
	eq(t, contains(r, 2), true)
	eq(t, contains(r, 3), true)
}

func TestHashSetCanLoopUsingFor(t *testing.T) {
	set := From[int](1, 2, 3)
	r := make([]int, 0, len(set))
	for item := range set {
		r = append(r, item)
	}	
	eq(t, contains(r, 1), true)
	eq(t, contains(r, 2), true)
	eq(t, contains(r, 3), true)
}
