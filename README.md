# golang-set
set implementation for Go

### Background
The missing set collection for the Go language.

### Usage
#### Int64HashSet
Set of int64 implemented with map

#### GenericHashSet[T comparable]
Set of generics implemented with map

#### HashSet
Set of interface{} implemented with map

#### Int64LinkHashSet
Ordered set of Int64 implemented with map

#### GenericLinkHashSet
Ordered set of generics implemented using map

#### LinkHashSet
Ordered set of interface{} implemented with map

### Performance
#### Insert
```
goos: darwin
goarch: arm64
pkg: set/set
BenchmarkInt64HashSetInsert-8           12088309               124.1 ns/op            34 B/op          0 allocs/op
BenchmarkGenericHashSetInsert-8         14383982               133.2 ns/op            55 B/op          0 allocs/op
BenchmarkHashSetInsert-8                 7372069               248.7 ns/op           101 B/op          1 allocs/op
BenchmarkGenericLinkHashSetInsert-8      9615259               191.0 ns/op            57 B/op          1 allocs/op
BenchmarkInt64LinkHashSetInsert-8       10234540               189.4 ns/op            55 B/op          1 allocs/op
BenchmarkLinkHashSetInsert-8             5307361               284.6 ns/op            96 B/op          2 allocs/op
PASS
ok      set/set 12.556s
```

#### List
```
goos: darwin
goarch: arm64
pkg: set/set
BenchmarkInt64HashSetList-8             1000000000               0.009527 ns/op        0 B/op          0 allocs/op
BenchmarkGenericHashSetList-8           1000000000               0.009524 ns/op        0 B/op          0 allocs/op
BenchmarkHashSetList-8                  1000000000               0.01232 ns/op         0 B/op          0 allocs/op
BenchmarkInt64LinkHashSetList-8         1000000000               0.001755 ns/op        0 B/op          0 allocs/op
BenchmarkGenericLinkHashSetList-8       1000000000               0.001776 ns/op        0 B/op          0 allocs/op
BenchmarkLinkHashSetList-8              1000000000               0.002497 ns/op        0 B/op          0 allocs/op
PASS
ok      set/set 7.174s
```
#### Delete
```
goos: darwin
goarch: arm64
pkg: set/set
BenchmarkInt64HashSetDelete-8           1000000000               0.0000059 ns/op               0 B/op          0 allocs/op
BenchmarkGenericHashSetDelete-8         1000000000               0.0000055 ns/op               0 B/op          0 allocs/op
BenchmarkHashSetDelete-8                1000000000               0.0000170 ns/op               0 B/op          0 allocs/op
BenchmarkInt64LinkHashSetDelete-8       1000000000               0.0002823 ns/op               0 B/op          0 allocs/op
BenchmarkGenericLinkHashSetDelete-8     1000000000               0.0002453 ns/op               0 B/op          0 allocs/op
BenchmarkLinkHashSetDelete-8            1000000000               0.0006696 ns/op               0 B/op          0 allocs/op
PASS
ok      set/set 0.616s
```

At all, the performance of Int64HashSet and GenericHashSet is similar and better.         
HashSet and LinkHashSet use reflection, the performance is reduced.             
In addition to using Map, LinkHashSet, Int64LinkHashSet and GenericLinkHashSet also uses one-way linked list to ensure the order of insertion.          

### Usage suggestions
For Golang >= 1.8, the interface{} is recommended, although there is a performance loss, it can reduce repeated code.    
For Golang < 1.8, the generics is recommended.    



