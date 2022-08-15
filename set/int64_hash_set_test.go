package set

import "testing"

func TestInt64HashSet(t *testing.T) {
	int64HashSet := NewInt64HashSet(8)
	int64HashSet.Insert(int64(111))
	if int64HashSet.Size() != 1 {
		t.Errorf("Expected size=1, but %d", int64HashSet.Size())
	}
	if !int64HashSet.Has(int64(111)) {
		t.Errorf("Expected contain `111`, but %#v", int64HashSet)
	}
	int64HashSet.Delete(111)
	if int64HashSet.Has(111) {
		t.Errorf("Unexpected contain `111`, but %#v", int64HashSet)
	}

	int64HashSet.Insert(int64(222), int64(333), int64(444), int64(555))
	if !int64HashSet.Has(int64(333)) {
		t.Errorf("Expected contain `333`, but %#v", int64HashSet)
	}
	int64HashSet.Delete(int64(111), int64(222), int64(555))
	if int64HashSet.Has(int64(222)) {
		t.Errorf("Unexpected contain `222`, but %#v", int64HashSet)
	}

	int64HashSet.Clear()
	if int64HashSet.Size() != 0 {
		t.Errorf("Expected size=0, but %#v", int64HashSet.Size())
	}

	int64HashSet.Insert(int64(1111), int64(2222), int64(3333))
	if len(int64HashSet.List()) != 3 {
		t.Errorf("Expected len(List)=3, but %#v", len(int64HashSet.List()))
	}
}
