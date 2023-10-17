package hashset

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/hashmap"
	"github.com/kkkunny/stl/container/iterator"
)

// HashSet 哈希set
type HashSet[T any] struct {
    data hashmap.HashMap[T, struct{}]
}

func NewHashSet[T any]() HashSet[T] {
    return HashSet[T]{data: hashmap.NewHashMap[T, struct{}]()}
}

func NewHashSetWithCapacity[T any](cap uint) HashSet[T] {
    return HashSet[T]{data: hashmap.NewHashMapWithCapacity[T, struct{}](cap)}
}

func NewHashSetWith[T any](vs ...T) HashSet[T] {
    self := NewHashSetWithCapacity[T](uint(len(vs)))
    for _, v := range vs {
        self.Push(v)
    }
    return self
}

func (_ HashSet[T]) NewWithIterator(iter iterator.Iterator[T]) HashSet[T] {
    self := NewHashSetWithCapacity[T](iter.Length())
    for iter.Next() {
        self.Push(iter.Value())
    }
    return self
}

func (self HashSet[T]) Length() uint {
    return self.data.Length()
}

func (self HashSet[T]) Equal(dst HashSet[T]) bool {
    return self.data.Equal(dst.data)
}

func (self *HashSet[T]) Push(v T) bool {
    exist := self.Contain(v)
    self.data.Set(v, struct{}{})
    return !exist
}

func (self HashSet[T]) Contain(v T) bool {
    return self.data.ContainKey(v)
}

func (self *HashSet[T]) Remove(v T) {
    self.data.Remove(v)
}

func (self HashSet[T]) String() string {
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

func (self HashSet[T]) Debug(prefix uint) string {
    var buf strings.Builder
    buf.WriteString("hashset{")
    var i int
    for iter := self.Iterator(); iter.Next(); {
        buf.WriteString(stlbasic.Debug(iter.Value(), prefix))
        if iter.HasNext() {
            buf.WriteString(", ")
        }
        i++
    }
    buf.WriteByte('}')
    return buf.String()
}

func (self HashSet[T]) Clone() HashSet[T] {
    hs := NewHashSetWithCapacity[T](self.Length())
    for iter := self.Iterator(); iter.Next(); {
        hs.Push(iter.Value())
    }
    return hs
}

func (self *HashSet[T]) Clear() {
    self.data.Clear()
}

func (self HashSet[T]) Empty() bool {
    return self.data.Empty()
}

func (self HashSet[T]) Iterator() iterator.Iterator[T] {
    return iterator.NewIterator[T](_NewIterator[T](&self))
}