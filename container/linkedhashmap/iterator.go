package linkedhashmap

import (
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/internal/list"
)

type _iterator[K, V any] struct {
	src  *LinkedHashMap[K, V]
	node *list.Element[pair.Pair[K, V]]
}

func _NewIterator[K, V any](src *LinkedHashMap[K, V]) *_iterator[K, V] {
	return &_iterator[K, V]{
		src:  src,
		node: nil,
	}
}

func (self _iterator[K, V]) Length() uint {
	return self.src.Length()
}

func (self *_iterator[K, V]) Next() bool {
	if self.node == nil{
		self.node = self.src.list.Front()
		return self.node != nil
	}else{
		if self.node.Next() == nil {
			return false
		}
		self.node = self.node.Next()
		return true
	}
}

func (self _iterator[K, V]) HasNext() bool {
	if self.node == nil{
		return self.Length() != 0
	}else{
		return self.node.Next() != nil
	}
}

func (self _iterator[K, V]) Value() pair.Pair[K, V] {
	return self.node.Value
}

func (self *_iterator[K, V]) Reset() {
	self.node = nil
}
