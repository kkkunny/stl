package treeset

import (
	"fmt"
	"strings"

	"github.com/HuKeping/rbtree"
	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/iterator"
)

type _TreeSetItem[T any] struct{
    value T
}

func (self _TreeSetItem[T]) Less(dst rbtree.Item) bool{
	return stlbasic.Order(self.value, dst.(_TreeSetItem[T]).value) < 0
}

// TreeSet 哈希set
type TreeSet[T any] struct {
	tree rbtree.Rbtree
}

func NewTreeSet[T any]() TreeSet[T] {
	return TreeSet[T]{tree: *rbtree.New()}
}

func NewTreeSetWith[T any](vs ...T) TreeSet[T] {
	self := NewTreeSet[T]()
	for _, v := range vs {
		self.Push(v)
	}
	return self
}

func (_ TreeSet[T]) NewWithIterator(iter iterator.Iterator[T]) TreeSet[T] {
	self := NewTreeSet[T]()
	for iter.Next() {
		self.Push(iter.Value())
	}
	return self
}

func (self TreeSet[T]) Length() uint {
	return self.tree.Len()
}

func (self TreeSet[T]) Equal(dst TreeSet[T]) bool {
    if self.Length() != dst.Length() {
		return false
	}

    equal := true
    dstIter := dst.Iterator()
    self.Iterator().Foreach(func(v T) bool{
        dstIter.Next()
        dv := dstIter.Value()

		if !stlbasic.Equal(v, dv) {
            equal = false
			return equal
		}
        return equal
    })
	return equal
}

func (self *TreeSet[T]) Push(v T) bool {
	if self.Contain(v){
		return false
	}
    self.tree.Insert(_TreeSetItem[T]{value: v})
	return true
}

func (self TreeSet[T]) Contain(v T) bool {
	return self.tree.Get(_TreeSetItem[T]{value: v}) != nil
}

func (self *TreeSet[T]) Remove(v T) {
	self.tree.Delete(_TreeSetItem[T]{value: v})
}

func (self TreeSet[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i int
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self TreeSet[T]) Clone() TreeSet[T] {
	hs := NewTreeSet[T]()
	for iter := self.Iterator(); iter.Next(); {
		hs.Push(iter.Value())
	}
	return hs
}

func (self *TreeSet[T]) Clear() {
	self.tree = *rbtree.New()
}

func (self TreeSet[T]) Empty() bool {
	return self.Length() == 0
}

func (self TreeSet[T]) Iterator() iterator.Iterator[T] {
	return iterator.NewIterator[T](_NewIterator[T](&self))
}

func (self TreeSet[T]) Back() T {
	return self.tree.Max().(_TreeSetItem[T]).value
}

func (self TreeSet[T]) Front() T {
	return self.tree.Min().(_TreeSetItem[T]).value
}
