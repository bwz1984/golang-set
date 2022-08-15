package set

import (
	"fmt"
	"strings"
)

type GenericHashSet[T comparable] map[T]Empty

var _ GenericSet[int] = (*GenericHashSet[int])(nil)

func NewGenericHashSet[T comparable](cap ...int) *GenericHashSet[T] {
	var set GenericHashSet[T]
	if len(cap) == 0 {
		set = make(GenericHashSet[T])
	} else {
		set = make(GenericHashSet[T], cap[0])
	}
	return &set
}

func (set *GenericHashSet[T]) Insert(items ...T) {
	for _, item := range items {
		(*set)[item] = Empty{}
	}
}

func (set *GenericHashSet[T]) Has(item T) bool {
	_, has := (*set)[item]
	return has
}

func (set *GenericHashSet[T]) Delete(items ...T) {
	for _, item := range items {
		delete(*set, item)
	}
}

func (set *GenericHashSet[T]) Clear() {
	for item := range *set {
		delete(*set, item)
	}
}

func (set *GenericHashSet[T]) Size() int {
	return len(*set)
}

func (set *GenericHashSet[T]) String() string {
	items := []string{}
	for item := range *set {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "[" + strings.Join(items, ",") + "]"
}

func (set *GenericHashSet[T]) List() []T {
	list := make([]T, 0, len(*set))
	for item := range *set {
		list = append(list, item)
	}
	return list
}
