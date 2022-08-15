package set

import (
	"fmt"
	"strings"
)

type int64ItemStruct struct {
	val  int64
	next *int64ItemStruct
}

type Int64LinkHashSet struct {
	hashSet map[int64]Empty
	header  *int64ItemStruct
	tailer  *int64ItemStruct
	cnt     int
}

func NewInt64LinkHashSet(args ...int) *Int64LinkHashSet {
	cap := 0
	if len(args) == 1 {
		cap = args[0]
	}
	header := &int64ItemStruct{next: nil}
	set := Int64LinkHashSet{
		hashSet: make(map[int64]Empty, cap),
		header:  header,
		tailer:  header,
		cnt:     0,
	}
	return &set
}

func (set *Int64LinkHashSet) Insert(items ...int64) {
	for _, item := range items {
		if _, ok := set.hashSet[item]; !ok {
			set.hashSet[item] = Empty{}
			newItem := &int64ItemStruct{val: item, next: nil}
			if set.cnt == 0 {
				set.header.next = newItem
			}
			set.tailer.next = newItem
			set.tailer = newItem
			set.cnt += 1
		}
	}
}

func (set *Int64LinkHashSet) Has(item int64) bool {
	_, has := set.hashSet[item]
	return has
}

func (set *Int64LinkHashSet) Delete(items ...int64) {
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

func (set *Int64LinkHashSet) Clear() {
	header := &int64ItemStruct{}
	set.hashSet = map[int64]Empty{}
	set.header = header
	set.tailer = header
	set.cnt = 0
}

func (set *Int64LinkHashSet) Size() int {
	return set.cnt
}

func (set *Int64LinkHashSet) List() []int64 {
	list := make([]int64, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		list = append(list, cur.val)
		cur = cur.next
	}
	return list
}

func (set *Int64LinkHashSet) String() string {
	items := []string{}
	cur := set.header.next
	for cur != nil {
		items = append(items, fmt.Sprintf("%v", cur.val))
		cur = cur.next
	}
	return "[" + strings.Join(items, ",") + "]"
}
