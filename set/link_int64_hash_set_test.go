package set

import (
	"reflect"
	"testing"
)

func TestLinkInt64HashSet(t *testing.T) {
	int64LinkhashSet := NewInt64LinkHashSet(8)
	int64LinkhashSet.Insert(int64(111))
	if int64LinkhashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", int64LinkhashSet.Size())
	}
	if !int64LinkhashSet.Has(int64(111)) {
		t.Errorf("Expected contain `111`, but %#v", int64LinkhashSet)
	}
	int64LinkhashSet.Delete(int64(111))
	if int64LinkhashSet.Has(int64(111)) {
		t.Errorf("Unexpected contain `111`, but %#v", int64LinkhashSet)
	}

	int64LinkhashSet.Insert(int64(222), int64(333), int64(444), int64(555))
	if !int64LinkhashSet.Has(int64(333)) {
		t.Errorf("Expected contain `333`, but %#v", int64LinkhashSet)
	}
	listInt64 := []int64{int64(222), int64(333), int64(444), int64(555)}
	if !reflect.DeepEqual(int64LinkhashSet.List(), listInt64) {
		t.Errorf("Expected euqal %#v but %#v", listInt64, int64LinkhashSet.List())
	}

	int64LinkhashSet.Delete(int64(111), int64(222), int64(555))
	if int64LinkhashSet.Has(int64(222)) {
		t.Errorf("Unexpected contain `222`, but %#v", int64LinkhashSet)
	}
	listInt64 = []int64{int64(333), int64(444)}
	if !reflect.DeepEqual(int64LinkhashSet.List(), listInt64) {
		t.Errorf("Expected euqal %#v but %#v", listInt64, int64LinkhashSet.List())
	}

	int64LinkhashSet.Clear()
	if int64LinkhashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", int64LinkhashSet.Size())
	}

	int64LinkhashSet.Insert(int64(1111), int64(2222), int64(3333))
	if len(int64LinkhashSet.List()) != 3 {
		t.Errorf("Expected len(List)=3, but %#v", len(int64LinkhashSet.List()))
	}
}
