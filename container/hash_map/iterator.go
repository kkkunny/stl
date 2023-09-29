package hashmap

import "github.com/kkkunny/stl/container/pair"

type _iterator[K comparable, V any] struct {
	src  *HashMap[K, V]
	keys []K
	next uint
}

func _NewIterator[K comparable, V any](src *HashMap[K, V]) *_iterator[K, V] {
	keys := make([]K, src.Length())
	var i int
	for k := range src.data {
		keys[i] = k
		i++
	}
	return &_iterator[K, V]{
		src:  src,
		keys: keys,
		next: 0,
	}
}

func (self *_iterator[K, V]) Length() uint {
	return self.src.Length()
}

func (self *_iterator[K, V]) Next() bool {
	if self.next >= self.Length() {
		return false
	}
	self.next++
	return true
}

func (self _iterator[K, V]) Value() pair.Pair[K, V] {
	key := self.keys[self.next-1]
	val := self.src.Get(key)
	return pair.NewPair[K, V](key, val)
}

func (self *_iterator[K, V]) Reset() {
	self.next = 0
}
