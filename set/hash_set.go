package set

import (
	"fmt"
	"strings"
)

type HashSet map[interface{}]Empty

var _ Set = (HashSet)(nil)

func NewHashSet(cap ...int) HashSet {
	var set HashSet
	if len(cap) == 0 {
		set = make(HashSet)
	} else {
		set = make(HashSet, cap[0])
	}
	return set
}

func (set HashSet) Insert(items ...interface{}) {
	for _, item := range items {
		set[item] = Empty{}
	}
}

func (set HashSet) Has(item interface{}) bool {
	_, has := set[item]
	return has
}

func (set HashSet) Delete(items ...interface{}) {
	for _, item := range items {
		delete(set, item)
	}
}

func (set HashSet) Clear() {
	for item := range set {
		delete(set, item)
	}
}

func (set HashSet) Size() int {
	return len(set)
}

func (set HashSet) String() string {
	items := []string{}
	for item := range set {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "[" + strings.Join(items, ",") + "]"
}

func (set HashSet) Diff(set2 HashSet) HashSet {
	res := NewHashSet()
	for item := range set {
		if !set2.Has(item) {
			res.Insert(item)
		}
	}
	return res
}

func (set HashSet) List() []interface{} {
	list := make([]interface{}, 0, len(set))
	for item := range set {
		list = append(list, item)
	}
	return list
}

func (set HashSet) ListByte() []byte {
	list := make([]byte, 0, len(set))
	for item := range set {
		if val, ok := item.(byte); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListInt8() []int8 {
	list := make([]int8, 0, len(set))
	for item := range set {
		if val, ok := item.(int8); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListUint8() []uint8 {
	list := make([]uint8, 0, len(set))
	for item := range set {
		if val, ok := item.(uint8); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListInt() []int {
	list := make([]int, 0, len(set))
	for item := range set {
		if val, ok := item.(int); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListUint() []uint {
	list := make([]uint, 0, len(set))
	for item := range set {
		if val, ok := item.(uint); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListInt32() []int32 {
	list := make([]int32, 0, len(set))
	for item := range set {
		if val, ok := item.(int32); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListUint32() []uint32 {
	list := make([]uint32, 0, len(set))
	for item := range set {
		if val, ok := item.(uint32); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListInt64() []int64 {
	list := make([]int64, 0, len(set))
	for item := range set {
		if val, ok := item.(int64); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListUint64() []uint64 {
	list := make([]uint64, 0, len(set))
	for item := range set {
		if val, ok := item.(uint64); ok {
			list = append(list, val)
		}
	}
	return list
}

func (set HashSet) ListString() []string {
	list := make([]string, 0, len(set))
	for item := range set {
		if val, ok := item.(string); ok {
			list = append(list, val)
		}
	}
	return list
}
