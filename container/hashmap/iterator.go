package hashmap

import "github.com/kkkunny/stl/container/pair"

type _iterator[K, V any] struct {
    src  *HashMap[K, V]
    pairs []pair.Pair[K, V]
    next uint
}

func _NewIterator[K, V any](src *HashMap[K, V]) *_iterator[K, V] {
    pairs := make([]pair.Pair[K, V], src.Length())
    var i int
    for _, pair := range src.data {
        pairs[i] = pair
        i++
    }
    return &_iterator[K, V]{
        src:  src,
        pairs: pairs,
        next: 0,
    }
}

func (self _iterator[K, V]) Length() uint {
    return self.src.Length()
}

func (self *_iterator[K, V]) Next() bool {
    if self.next >= self.Length() {
        return false
    }
    self.next++
    return true
}

func (self _iterator[K, V]) HasNext() bool {
    return self.next < self.Length()
}

func (self _iterator[K, V]) Value() pair.Pair[K, V] {
    return self.pairs[self.next-1]
}

func (self *_iterator[K, V]) Reset() {
    self.next = 0
}
