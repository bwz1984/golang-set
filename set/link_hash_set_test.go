package set

import (
	"reflect"
	"testing"
)

func TestStringLinkHashSet(t *testing.T) {
	linkhashSet := NewLinkHashSet(8)
	linkhashSet.Insert("aa")
	if linkhashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", linkhashSet.Size())
	}
	if !linkhashSet.Has("aa") {
		t.Errorf("Expected contain `aa`, but %#v", linkhashSet)
	}
	linkhashSet.Delete("aa")
	if linkhashSet.Has("aa") {
		t.Errorf("Unexpected contain `aa`, but %#v", linkhashSet)
	}

	linkhashSet.Insert("bb", "cc", "dd", "ee")
	if !linkhashSet.Has("cc") {
		t.Errorf("Expected contain `cc`, but %#v", linkhashSet)
	}
	strs := []string{"bb", "cc", "dd", "ee"}
	if !reflect.DeepEqual(linkhashSet.ListString(), strs) {
		t.Errorf("Expected euqal %#v but %#v", strs, linkhashSet.ListString())
	}

	linkhashSet.Delete("aa", "dd", "cc")
	if linkhashSet.Has("dd") {
		t.Errorf("Unexpected contain `dd`, but %#v", linkhashSet)
	}
	strs = []string{"bb", "ee"}
	if !reflect.DeepEqual(linkhashSet.ListString(), strs) {
		t.Errorf("Expected euqal %#v but %#v", strs, linkhashSet.ListString())
	}

	linkhashSet.Clear()
	if linkhashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", linkhashSet.Size())
	}

	linkhashSet.Insert("11", "22", "33")
	if len(linkhashSet.ListString()) != 3 {
		t.Errorf("Expected len(ListString)=3, but %#v", len(linkhashSet.ListString()))
	}
}

func TestInt64LinkHashSet(t *testing.T) {
	linkhashSet := NewLinkHashSet(8)
	linkhashSet.Insert(int64(111))
	if linkhashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", linkhashSet.Size())
	}
	if !linkhashSet.Has(int64(111)) {
		t.Errorf("Expected contain `111`, but %#v", linkhashSet)
	}
	linkhashSet.Delete(int64(111))
	if linkhashSet.Has(int64(111)) {
		t.Errorf("Unexpected contain `111`, but %#v", linkhashSet)
	}

	linkhashSet.Insert(int64(222), int64(333), int64(444), int64(555))
	if !linkhashSet.Has(int64(333)) {
		t.Errorf("Expected contain `333`, but %#v", linkhashSet)
	}
	listInt64 := []int64{int64(222), int64(333), int64(444), int64(555)}
	if !reflect.DeepEqual(linkhashSet.ListInt64(), listInt64) {
		t.Errorf("Expected euqal %#v but %#v", listInt64, linkhashSet.ListInt64())
	}

	linkhashSet.Delete(int64(111), int64(222), int64(555))
	if linkhashSet.Has(int64(222)) {
		t.Errorf("Unexpected contain `222`, but %#v", linkhashSet)
	}
	listInt64 = []int64{int64(333), int64(444)}
	if !reflect.DeepEqual(linkhashSet.ListInt64(), listInt64) {
		t.Errorf("Expected euqal %#v but %#v", listInt64, linkhashSet.ListInt64())
	}

	linkhashSet.Clear()
	if linkhashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", linkhashSet.Size())
	}

	linkhashSet.Insert(int64(1111), int64(2222), int64(3333))
	if len(linkhashSet.ListInt64()) != 3 {
		t.Errorf("Expected len(ListInt64)=3, but %#v", len(linkhashSet.ListInt64()))
	}
}
