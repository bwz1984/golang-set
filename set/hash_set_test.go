package set

import (
	"reflect"
	"testing"
)

func TestHashSetString(t *testing.T) {
	hashSet := NewHashSet(8)
	hashSet.Insert("aa")
	if hashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", hashSet.Size())
	}
	if !hashSet.Has("aa") {
		t.Errorf("Expected contain `aa`, but %#v", hashSet)
	}
	hashSet.Delete("aa")
	if hashSet.Has("aa") {
		t.Errorf("Unexpected contain `aa`, but %#v", hashSet)
	}

	hashSet.Insert("bb", "cc", "dd", "ee")
	if !hashSet.Has("cc") {
		t.Errorf("Expected contain `cc`, but %#v", hashSet)
	}
	hashSet.Delete("aa", "dd", "cc")
	if hashSet.Has("dd") {
		t.Errorf("Unexpected contain `dd`, but %#v", hashSet)
	}

	hashSet.Clear()
	if hashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", hashSet.Size())
	}

	hashSet.Insert("11", "22", "33")
	if len(hashSet.ListString()) != 3 {
		t.Errorf("Expected len(ListString)=3, but %#v", len(hashSet.ListString()))
	}
}

func TestHashSetInt64(t *testing.T) {
	hashSet := NewHashSet(8)
	hashSet.Insert(int64(111))
	if hashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", hashSet.Size())
	}
	if !hashSet.Has(int64(111)) {
		t.Errorf("Expected contain `111`, but %#v", hashSet)
	}
	hashSet.Delete(111)
	if hashSet.Has(111) {
		t.Errorf("Unexpected contain `111`, but %#v", hashSet)
	}

	hashSet.Insert(int64(222), int64(333), int64(444), int64(555))
	if !hashSet.Has(int64(333)) {
		t.Errorf("Expected contain `333`, but %#v", hashSet)
	}
	hashSet.Delete(int64(111), int64(222), int64(555))
	if hashSet.Has(int64(222)) {
		t.Errorf("Unexpected contain `222`, but %#v", hashSet)
	}

	hashSet.Clear()
	if hashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", hashSet.Size())
	}

	hashSet.Insert(int64(1111), int64(2222), int64(3333))
	if len(hashSet.ListInt64()) != 3 {
		t.Errorf("Expected len(ListInt64)=3, but %#v", len(hashSet.ListInt64()))
	}

}

func TestDiff(t *testing.T) {
	hashSet := NewHashSet()
	hashSet.Insert(1111, 2222, 3333)
	hashSet2 := NewHashSet()
	hashSet2.Insert(1111)
	diffSet := hashSet.Diff(hashSet2)
	diffSetE := NewHashSet()
	diffSetE.Insert(2222, 3333)
	if !reflect.DeepEqual(diffSet, diffSetE) {
		t.Errorf("Expected %#v, but %#v", diffSet, diffSetE)
	}
}
