package set

import "testing"

func TestGenericHashSetString(t *testing.T) {
	gHashSet := NewGenericHashSet[string]()
	gHashSet.Insert("aa")
	if gHashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", gHashSet.Size())
	}
	if !gHashSet.Has("aa") {
		t.Errorf("Expected contain `aa`, but %#v", gHashSet)
	}
	gHashSet.Delete("aa")
	if gHashSet.Has("aa") {
		t.Errorf("Unexpected contain `aa`, but %#v", gHashSet)
	}

	gHashSet.Insert("bb", "cc", "dd", "ee")
	if !gHashSet.Has("cc") {
		t.Errorf("Expected contain `cc`, but %#v", gHashSet)
	}
	gHashSet.Delete("aa", "dd", "cc")
	if gHashSet.Has("dd") {
		t.Errorf("Unexpected contain `dd`, but %#v", gHashSet)
	}

	gHashSet.Clear()
	if gHashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", gHashSet.Size())
	}

	gHashSet.Insert("11", "22", "33")
	if len(gHashSet.List()) != 3 {
		t.Errorf("Expected len(ListString)=3, but %#v", len(gHashSet.List()))
	}
}

func TestGenericHashSetInt64(t *testing.T) {
	gHashSet := NewGenericHashSet[int64]()
	gHashSet.Insert(int64(111))
	if gHashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", gHashSet.Size())
	}
	if !gHashSet.Has(int64(111)) {
		t.Errorf("Expected contain `111`, but %#v", gHashSet)
	}
	gHashSet.Delete(111)
	if gHashSet.Has(111) {
		t.Errorf("Unexpected contain `111`, but %#v", gHashSet)
	}

	gHashSet.Insert(int64(222), int64(333), int64(444), int64(555))
	if !gHashSet.Has(int64(333)) {
		t.Errorf("Expected contain `333`, but %#v", gHashSet)
	}
	gHashSet.Delete(int64(111), int64(222), int64(555))
	if gHashSet.Has(int64(222)) {
		t.Errorf("Unexpected contain `222`, but %#v", gHashSet)
	}

	gHashSet.Clear()
	if gHashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", gHashSet.Size())
	}

	gHashSet.Insert(int64(1111), int64(2222), int64(3333))
	if len(gHashSet.List()) != 3 {
		t.Errorf("Expected len(ListInt64)=3, but %#v", len(gHashSet.List()))
	}
}
