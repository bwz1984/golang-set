# golang-set
set implementation for Go

### Background
The missing set collection for the Go language.

### Usage
#### Int64HashSet
Set of Int64 implemented with map

#### HashSet
Set of interface{} implemented with map

#### LinkHashSet
Ordered set of interfac{} implemented using map

### Performance
#### Insert
```
goos: darwin
goarch: arm64
pkg: set/set
BenchmarkInt64HashSetInsert-8            9821091               119.3 ns/op            40 B/op          0 allocs/op
BenchmarkHashSetInsert-8                 7083038               231.9 ns/op           104 B/op          1 allocs/op
BenchmarkLinkHashSetInsert-8             4998378               277.4 ns/op           100 B/op          2 allocs/op
PASS
ok      set/set 6.506s
```
#### List
```
goos: darwin
goarch: arm64
pkg: set/set
BenchmarkInt64HashSetList-8     1000000000               0.009093 ns/op        0 B/op          0 allocs/op
BenchmarkHashSetList-8          1000000000               0.009796 ns/op        0 B/op          0 allocs/op
BenchmarkLinkHashSetList-8      1000000000               0.002174 ns/op        0 B/op          0 allocs/op
PASS
ok      set/set 3.680s
```
#### Delete
```
goos: darwin
goarch: arm64
pkg: set/set
BenchmarkInt64HashSetDelete-8           1000000000               0.0000060 ns/op               0 B/op          0 allocs/op
BenchmarkHashSetDelete-8                1000000000               0.0000120 ns/op               0 B/op          0 allocs/op
BenchmarkLinkHashSetDelete-8            1000000000               0.0009641 ns/op               0 B/op          0 allocs/op
PASS
ok      set/set 0.285s
```
