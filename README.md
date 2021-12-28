


<p align="center">
<img align="center" width="500px" src="gollections.png">
</p> 
<p align="center">
Implementation of missing colections for Golang using Generics. <br> 
Free of dependencies.  <br />
Curently in <b>early development</b> phase.
</p>  

## Requirements

Go **1.18+**

## Install

```console
$ go get -u github.com/ferreiravinicius/gollections
```

## Set

Set is a collection that contains no duplicate elements.  
All implementations follows the set.Set interface.

## Hash Set

Implemention of set backed by a hash table (Go builtin Map).  
Order of insertion is not guaranteed.  
This implementation is not synchronized.

#### **API**

Importing
```golang
import "github.com/ferreiravinicius/goset/hashset"
```

Create a new empty set
```golang
set := hashset.New[int]() 
```

Create a new set initialized
```golang
set := hashset.From(1, 2, 3) 

items := []int{4, 5, 6}
set := hashset.From(items...) 
```

Create a new set with initial capacity 
```golang
set := hashset.WithCapacity[int](100) 
```

Iterating over the set using builtin loop

```golang
set := hashset.From(1, 2, 3)
for item := range set {
  // foo(item)
}
```

Iterating over the set using `ForEach`
```golang
set := hashset.From(1, 2, 3) 
set.ForEach(func(item int) { 
 // foo(item) 
})
```
