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
	//fmt.Println(Int64HashSet.List())
}

func BenchmarkHashSetInsert(b *testing.B) {
	hashSet := NewHashSet()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		hashSet.Insert(rand.Int63())
	}
	//fmt.Println(hashSet.List())
}

func BenchmarkLinkHashSetInsert(b *testing.B) {
	linkHashSet := NewLinkHashSet()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		linkHashSet.Insert(rand.Int63())
	}
	//fmt.Println(LinkHashSet.List())
}

func BenchmarkInt64HashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	intHashSet := NewInt64HashSet(0)
	for i := 0; i < 1024*1024; i++ {
		intHashSet.Insert(rand.Int63())
	}
	//b.ResetTimer()
	_ = intHashSet.List()
}

func BenchmarkHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	hashSet := NewHashSet(0)
	for i := 0; i < 1024*1024; i++ {
		hashSet.Insert(rand.Int())
	}
	//b.ResetTimer()
	_ = hashSet.ListInt()
}

func BenchmarkLinkHashSetList(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	linkHashSet := NewLinkHashSet(0)
	for i := 0; i < 1024*1024; i++ {
		linkHashSet.Insert(rand.Int())
	}
	//b.ResetTimer()
	_ = linkHashSet.ListInt()
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

func BenchmarkHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	hashSet := NewHashSet(0)
	for i := 0; i < 1024; i++ {
		hashSet.Insert(rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		hashSet.Delete(rand.Int())
	}
}

func BenchmarkLinkHashSetDelete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	linkHashSet := NewLinkHashSet(0)
	for i := 0; i < 1024; i++ {
		linkHashSet.Insert(rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < 128; i++ {
		linkHashSet.Delete(rand.Int())
	}
}
