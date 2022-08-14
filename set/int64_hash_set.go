package set

import (
	"fmt"
	"strings"
)

type Int64HashSet map[int64]Empty

func NewInt64HashSet(cap ...int) Int64HashSet {
	var set Int64HashSet
	if len(cap) == 0 {
		set = make(Int64HashSet)
	} else {
		set = make(Int64HashSet, cap[0])
	}
	return set
}

func (set Int64HashSet) Insert(items ...int64) {
	for _, item := range items {
		set[item] = Empty{}
	}
}

func (set Int64HashSet) Has(item int64) bool {
	_, has := set[item]
	return has
}

func (set Int64HashSet) Delete(items ...int64) {
	for _, item := range items {
		delete(set, item)
	}
}

func (set Int64HashSet) Clear() {
	for item := range set {
		delete(set, item)
	}
}

func (set Int64HashSet) Size() int {
	return len(set)
}

func (set Int64HashSet) List() []int64 {
	list := make([]int64, 0, len(set))
	for item := range set {
		list = append(list, item)
	}
	return list
}

func (set Int64HashSet) String() string {
	items := make([]string, 0, len(set))
	for item := range set {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "[" + strings.Join(items, ",") + "]"
}
