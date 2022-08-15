package set

import (
	"fmt"
	"strings"
)

type genericItemStruct[T comparable] struct {
	val  T
	next *genericItemStruct[T]
}

type GenericLinkHashSet[T comparable] struct {
	hashSet map[T]Empty
	header  *genericItemStruct[T]
	tailer  *genericItemStruct[T]
	cnt     int
}

var _ GenericSet[int] = (*GenericLinkHashSet[int])(nil)

func NewGenericLinkHashSet[T comparable](args ...int) *GenericLinkHashSet[T] {
	cap := 0
	if len(args) == 1 {
		cap = args[0]
	}
	header := &genericItemStruct[T]{next: nil}
	set := GenericLinkHashSet[T]{
		hashSet: make(map[T]Empty, cap),
		header:  header,
		tailer:  header,
		cnt:     0,
	}
	return &set
}

func (set *GenericLinkHashSet[T]) Insert(items ...T) {
	for _, item := range items {
		if _, ok := set.hashSet[item]; !ok {
			set.hashSet[item] = Empty{}
			newItem := &genericItemStruct[T]{val: item, next: nil}
			if set.cnt == 0 {
				set.header.next = newItem
			}
			set.tailer.next = newItem
			set.tailer = newItem
			set.cnt += 1
		}
	}
}

func (set *GenericLinkHashSet[T]) Has(item T) bool {
	_, has := set.hashSet[item]
	return has
}

func (set *GenericLinkHashSet[T]) Delete(items ...T) {
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

func (set *GenericLinkHashSet[T]) Clear() {
	header := &genericItemStruct[T]{}
	set.hashSet = map[T]Empty{}
	set.header = header
	set.tailer = header
	set.cnt = 0
}

func (set *GenericLinkHashSet[T]) Size() int {
	return set.cnt
}

func (set *GenericLinkHashSet[T]) List() []T {
	list := make([]T, 0, set.cnt)
	cur := set.header.next
	for cur != nil {
		list = append(list, cur.val)
		cur = cur.next
	}
	return list
}

func (set *GenericLinkHashSet[T]) String() string {
	items := []string{}
	cur := set.header.next
	for cur != nil {
		items = append(items, fmt.Sprintf("%v", cur.val))
		cur = cur.next
	}
	return "[" + strings.Join(items, ",") + "]"
}
