package set

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkInt64HashSetInsert(b *testing.B) {
	intHashSet := NewInt64HashSet()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		intHashSet.Insert(rand.Int63())
	}
}

func BenchmarkGenericHashSetInsert(b *testing.B) {
	gHashSet := NewGenericHashSet[int64]()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		gHashSet.Insert(rand.Int63())
	}
}

func BenchmarkHashSetInsert(b *testing.B) {
	hashSet := NewHashSet()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		hashSet.Insert(rand.Int63())
	}
}

func BenchmarkGenericLinkHashSetInsert(b *testing.B) {
	linkGenericHashSet := NewGenericLinkHashSet[int64]()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		linkGenericHashSet.Insert(rand.Int63())
	}
}

func BenchmarkInt64LinkHashSetInsert(b *testing.B) {
	int64LinkHashSet := NewInt64LinkHashSet()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		int64LinkHashSet.Insert(rand.Int63())
	}
}

func BenchmarkLinkHashSetInsert(b *testing.B) {
	linkHashSet := NewLinkHashSet()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		linkHashSet.Insert(rand.Int63())
	}
}

func BenchmarkInt64HashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	intHashSet := NewInt64HashSet(0)
	for i := 0; i < 1024*1024; i++ {
		intHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	_ = intHashSet.List()
}

func BenchmarkGenericHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	hashSet := NewGenericHashSet[int64]()
	for i := 0; i < 1024*1024; i++ {
		hashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	_ = hashSet.List()
}

func BenchmarkHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	hashSet := NewHashSet(0)
	for i := 0; i < 1024*1024; i++ {
		hashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	_ = hashSet.ListInt64()
}

func BenchmarkInt64LinkHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	int64LinkHashSet := NewInt64LinkHashSet(0)
	for i := 0; i < 1024*1024; i++ {
		int64LinkHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	_ = int64LinkHashSet.List()
}

func BenchmarkGenericLinkHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	gLinkHashSet := NewGenericLinkHashSet[int64](0)
	for i := 0; i < 1024*1024; i++ {
		gLinkHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	_ = gLinkHashSet.List()
}

func BenchmarkLinkHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	linkHashSet := NewLinkHashSet(0)
	for i := 0; i < 1024*1024; i++ {
		linkHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	_ = linkHashSet.ListInt64()
}

func BenchmarkInt64HashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	intHashSet := NewInt64HashSet(0)
	for i := 0; i < 1024; i++ {
		intHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		intHashSet.Delete(rand.Int63())
	}
}

func BenchmarkGenericHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	hashSet := NewGenericHashSet[int64]()
	for i := 0; i < 1024; i++ {
		hashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		hashSet.Delete(rand.Int63())
	}
}

func BenchmarkHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	hashSet := NewHashSet(0)
	for i := 0; i < 1024; i++ {
		hashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		hashSet.Delete(rand.Int63())
	}
}

func BenchmarkInt64LinkHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	int64LinkHashSet := NewInt64LinkHashSet(0)
	for i := 0; i < 1024; i++ {
		int64LinkHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		int64LinkHashSet.Delete(rand.Int63())
	}
}

func BenchmarkGenericLinkHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	gLinkHashSet := NewGenericLinkHashSet[int64](0)
	for i := 0; i < 1024; i++ {
		gLinkHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		gLinkHashSet.Delete(rand.Int63())
	}
}

func BenchmarkLinkHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	linkHashSet := NewLinkHashSet(0)
	for i := 0; i < 1024; i++ {
		linkHashSet.Insert(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		linkHashSet.Delete(rand.Int63())
	}
}
