package hashmap

// 桶
type bucket[K, V any] struct {
	Used  bool
	Hash  uint64
	Key   K
	Value V
}

// HashMap 哈希表
type HashMap[K, V any] struct {
	buckets *[]bucket[K, V]
	length  uint
}

func NewHashMap[K, V any]() HashMap[K, V] {
	return NewHashMapWithCapacity[K, V](initialBucketCapacity)
}

func NewHashMapWithCapacity[K, V any](cap uint) HashMap[K, V] {
	if cap == 0 {
		cap = 1
	}
	var buckets []bucket[K, V]
	if cap < initialBucketCapacity {
		buckets = make([]bucket[K, V], initialBucketCapacity)
	} else {
		buckets = make([]bucket[K, V], cap)
	}
	return HashMap[K, V]{buckets: &buckets}
}

func NewHashMapWith[K, V any](vs ...any) HashMap[K, V] {
	self := NewHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self HashMap[K, V]) Default() HashMap[K, V] {
	return NewHashMap[K, V]()
}
