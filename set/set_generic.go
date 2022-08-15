package set

type GenericSet[T comparable] interface {
	Insert(key ...T)
	Has(key T) bool
	Delete(key ...T)
	Clear()
	Size() int
	List() []T
	String() string
}
