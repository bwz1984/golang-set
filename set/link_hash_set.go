package set

import (
	"fmt"
	"strings"
)

type itemStruct struct {
	val  interface{}
	next *itemStruct
}

type LinkHashSet struct {
	hashSet map[interface{}]Empty
	header  *itemStruct
	tailer  *itemStruct
	cnt     int
}

var _ Set = (*LinkHashSet)(nil)

func NewLinkHashSet(args ...int) *LinkHashSet {
	cap := 0
	if len(args) == 1 {
		cap = args[0]
	}
	header := &itemStruct{val: nil, next: nil}
	set := LinkHashSet{
		hashSet: make(map[interface{}]Empty, cap),
		header:  header,
		tailer:  header,
		cnt:     0,
	}
	return &set
}

func (set *LinkHashSet) Insert(items ...interface{}) {
	for _, item := range items {
		if _, ok := set.hashSet[item]; !ok {
			set.hashSet[item] = Empty{}
			newItem := &itemStruct{val: item, next: nil}
			if set.cnt == 0 {
				set.header.next = newItem
			}
			set.tailer.next = newItem
			set.tailer = newItem
			set.cnt += 1
		}
	}
}

func (set *LinkHashSet) Has(item interface{}) bool {
	_, has := set.hashSet[item]
	return has
}

func (set *LinkHashSet) Delete(items ...interface{}) {
	for _, item := range items {
		delete(set.hashSet, item)
		cur := set.header
		for cur != nil && cur != set.tailer {
			if cur.next != nil && cur.next.val == item {
				if cur.next.val == set.tailer.val {
					set.tailer = cur
				}
				cur.next = cur.next.next
				set.cnt -= 1
				break
			}
			cur = cur.next
		}
	}
}

func (set *LinkHashSet) Clear() {
	header := &itemStruct{}
	set.hashSet = map[interface{}]Empty{}
	set.header = header
	set.tailer = header
	set.cnt = 0
}

func (set *LinkHashSet) Size() int {
	return set.cnt
}

func (set *LinkHashSet) String() string {
	items := []string{}
	cur := set.header.next
	for cur != nil {
		items = append(items, fmt.Sprintf("%v", cur.val))
		cur = cur.next
	}
	return "[" + strings.Join(items, ",") + "]"
}

func (set *LinkHashSet) Diff(set2 *LinkHashSet) *LinkHashSet {
	res := NewLinkHashSet()
	cur := set.header.next
	for cur != nil {
		if !set2.Has(cur.val) {
			res.Insert(cur.val)
		}
		cur = cur.next
	}
	return res
}

func (set *LinkHashSet) List() []interface{} {
	list := make([]interface{}, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		list = append(list, cur.val)
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListByte() []byte {
	list := make([]byte, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(byte); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListInt8() []int8 {
	list := make([]int8, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(int8); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListUint8() []uint8 {
	list := make([]uint8, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(uint8); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListInt() []int {
	list := make([]int, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(int); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListUint() []uint {
	list := make([]uint, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(uint); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListInt32() []int32 {
	list := make([]int32, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(int32); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListUint32() []uint32 {
	list := make([]uint32, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(uint32); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListInt64() []int64 {
	list := make([]int64, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(int64); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListUint64() []uint64 {
	list := make([]uint64, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(uint64); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}

func (set *LinkHashSet) ListString() []string {
	list := make([]string, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		if val, ok := cur.val.(string); ok {
			list = append(list, val)
		}
		cur = cur.next
	}
	return list
}
