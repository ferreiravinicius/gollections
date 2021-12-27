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

func TestHashSetNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Errorf("expected set not to be nil")
	}
	if s.hashMap == nil {
		t.Errorf("expected map inside of set not to be nil")
	}
	havingSize := New[int]()
	if havingSize.hashMap == nil {
		t.Errorf("expected map inside of set not to be nil")
	}
}

func TestHashSetAdd(t *testing.T) {
	s := New[int]()
	eq(t, s.Add(1), true)
	eq(t, len(s.hashMap), 1)
	eq(t, s.Add(1), false)
	eq(t, len(s.hashMap), 1)
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
	eq(t, len(s.hashMap), 0)

	eq(t, s.AddAll(1, 2, 3), true)
	eq(t, len(s.hashMap), 3)

	eq(t, s.AddAll(1, 2, 3, 4), true)
	eq(t, len(s.hashMap), 4)

	eq(t, s.AddAll(1, 2), false)
	eq(t, len(s.hashMap), 4)
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

	results := make(chan T, len(set.hashMap))
	cancel := make(chan bool, 1)
	exit := func() {
		cancel <- true
	}
	
	go func(results chan T, cancel chan bool) {
		defer close(results)
		defer close(cancel)
		for item := range set.hashMap {
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


