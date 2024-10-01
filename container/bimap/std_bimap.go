package bimap

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

type _StdBiMap[T, E comparable] struct {
	keys   hashmap.HashMap[T, E]
	values hashmap.HashMap[E, T]
}

func _NewStdBiMap[T, E comparable]() BiMap[T, E] {
	return &_StdBiMap[T, E]{
		keys:   hashmap.AnyWith[T, E](),
		values: hashmap.AnyWith[E, T](),
	}
}

func _NewStdBiMapWithCapacity[T, E comparable](cap uint) BiMap[T, E] {
	return &_StdBiMap[T, E]{
		keys:   hashmap.AnyWithCap[T, E](cap),
		values: hashmap.AnyWithCap[E, T](cap),
	}
}

func _NewStdBiMapWith[T, E comparable](vs ...any) BiMap[T, E] {
	self := _NewStdBiMapWithCapacity[T, E](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(T), vs[i+1].(E))
	}
	return self
}

func (self *_StdBiMap[T, E]) Capacity() uint {
	return self.keys.Capacity()
}

func (self *_StdBiMap[T, E]) Clone() BiMap[T, E] {
	return &_StdBiMap[T, E]{
		keys:   self.keys.Clone(),
		values: self.values.Clone(),
	}
}

func (self *_StdBiMap[T, E]) Equal(dst BiMap[T, E]) bool {
	return self.keys.Equal(dst.getKeyData())
}

func (_ *_StdBiMap[T, E]) NewWithIterator(iter stliter.Iterator[pair.Pair[T, E]]) any {
	self := _NewStdBiMapWithCapacity[T, E](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self *_StdBiMap[T, E]) Iterator() stliter.Iterator[pair.Pair[T, E]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_StdBiMap[T, E]) Length() uint {
	return self.keys.Length()
}

// Set 插入键值对
func (self *_StdBiMap[T, E]) Set(k T, v E) (T, E) {
	pv, pk := self.RemoveKey(k), self.RemoveValue(v)
	self.keys.Set(k, v)
	self.values.Set(v, k)
	return pk, pv
}

// GetValue 获取值
func (self *_StdBiMap[T, E]) GetValue(k T, defaultValue ...E) E {
	return self.keys.Get(k, defaultValue...)
}

// GetKey 获取键
func (self *_StdBiMap[T, E]) GetKey(v E, defaultKey ...T) T {
	return self.values.Get(v, defaultKey...)
}

// ContainKey 是否包含键
func (self *_StdBiMap[T, E]) ContainKey(k T) bool {
	return self.keys.ContainKey(k)
}

// ContainValue 是否包含值
func (self *_StdBiMap[T, E]) ContainValue(v E) bool {
	return self.values.ContainKey(v)
}

// RemoveKey 移除键
func (self *_StdBiMap[T, E]) RemoveKey(k T, defaultValue ...E) E {
	exist := self.ContainKey(k)
	pv := self.keys.Remove(k, defaultValue...)
	if exist {
		self.values.Remove(pv)
	}
	return pv
}

// RemoveValue 移除值
func (self *_StdBiMap[T, E]) RemoveValue(v E, defaultKey ...T) T {
	exist := self.ContainValue(v)
	pk := self.values.Remove(v, defaultKey...)
	if exist {
		self.keys.Remove(pk)
	}
	return pk
}

// Clear 清空
func (self *_StdBiMap[T, E]) Clear() {
	self.keys.Clear()
	self.values.Clear()
}

// Empty 是否为空
func (self *_StdBiMap[T, E]) Empty() bool {
	return self.keys.Empty()
}

// Keys 获取所有键
func (self *_StdBiMap[T, E]) Keys() []T {
	return self.keys.Keys()
}

// Values 获取所有值
func (self *_StdBiMap[T, E]) Values() []E {
	return self.values.Keys()
}

// KeyValues 获取所有键值对
func (self *_StdBiMap[T, E]) KeyValues() []pair.Pair[T, E] {
	return self.keys.KeyValues()
}

func (self *_StdBiMap[T, E]) String() string {
	var buf strings.Builder
	buf.WriteString("BiMap{")
	for i, p := range self.KeyValues() {
		buf.WriteString(fmt.Sprintf("%v", p.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", p.Second))
		if i < int(self.Length())-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_StdBiMap[T, E]) getKeyData() hashmap.HashMap[T, E] {
	return self.keys
}

func (self *_StdBiMap[T, E]) getValueData() hashmap.HashMap[E, T] {
	return self.values
}
