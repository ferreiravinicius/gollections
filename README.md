Implementation of missing colections for Golang using Generics.  
Free of dependencies.  
Curently in **early development** phase.

## Requirements

Go **1.18+**

## Install

```console
$ go get -u github.com/ferreiravinicius/gollections
```

# Set

Set is a collection that contains no duplicate elements.  
All implementations follows the set.Set interface.

## Hash Set

Implemention of set backed by a hash table (Go builtin Map).  
Order of insertion is not guaranteed.  
This implementation is not synchronized.


#### **API**

*Import package*
```golang
import "github.com/ferreiravinicius/goset/hashset"
```

Creating a new set  
```golang
set := hashset.New[int]() 
```

New set from items or slice
```golang
set := hashset.From(1, 2, 3) 

items := []int{4, 5, 6}
set := hashset.From(items...) 
```

New set with initial capacity (size).
```golang
set := hashset.WithCapacity[int](100) 
```

Iterating over using for loop

```golang
set := hashset.From(1, 2, 3)
for item := range set {
  // foo(item)
}
```


Iterating over a set using `ForEach`
```golang
set := hashset.From(1, 2, 3) 
set.ForEach(func(item int) { 
 // foo(item) 
})
```