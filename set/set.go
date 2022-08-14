package set

type Set interface {
	Insert(key ...interface{})
	Has(key interface{}) bool
	Delete(key ...interface{})
	Clear()
	Size() int
	List() []interface{}
	ListByte() []byte
	ListInt8() []int8
	ListUint8() []uint8
	ListInt() []int
	ListUint() []uint
	ListInt32() []int32
	ListUint32() []uint32
	ListInt64() []int64
	ListUint64() []uint64
	ListString() []string
	String() string
}
