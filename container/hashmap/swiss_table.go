package hashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kkkunny/swiss"

	stlbasic "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
)

var initSwissTableCapacity uint

func init() {
	initSwissTableCapacity = uint(swiss.NewMap[int, int](1).Capacity())
}

type _SwissTable[K comparable, V any] struct {
	data *swiss.Map[K, V]
}

func _NewSwissTable[K comparable, V any]() HashMap[K, V] {
	return _NewSwissTableWithCapacity[K, V](initSwissTableCapacity)
}

func _NewSwissTableWithCapacity[K comparable, V any](cap uint) HashMap[K, V] {
	return &_SwissTable[K, V]{data: swiss.NewMap[K, V](uint32(cap))}
}

func _NewSwissTableWith[K comparable, V any](vs ...any) HashMap[K, V] {
	self := _NewSwissTableWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self *_SwissTable[K, V]) Capacity() uint {
	return uint(self.data.Capacity() + self.data.Count())
}

func (self *_SwissTable[K, V]) Clone() any {
	newMap := _NewSwissTableWithCapacity[K, V](self.Capacity())
	self.data.Iter(func(k K, v V) bool {
		newMap.Set(k, v)
		return false
	})
	return newMap
}

func (self *_SwissTable[K, V]) Equal(dstObj any) (eq bool) {
	if dstObj == nil && self == nil {
		return true
	} else if dstObj == nil {
		return false
	}

	dst, ok := dstObj.(HashMap[K, V])
	if !ok {
		return false
	}

	if self.Length() != dst.Length() {
		return false
	}

	eq = true
	self.data.Iter(func(k K, sv V) bool {
		if !dst.Contain(k) || !stlbasic.Equal(sv, dst.Get(k)) {
			eq = false
			return true
		}
		return false
	})
	return eq
}

func (_ *_SwissTable[K, V]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[K, V]]) any {
	self := _NewSwissTableWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.Unpack())
	}
	return self
}

func (self *_SwissTable[K, V]) Iterator() stliter.Iterator[tuple.Tuple2[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_SwissTable[K, V]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := buf.WriteByte('{')
	if err != nil {
		return nil, err
	}
	var i int
	self.data.Iter(func(k K, v V) bool {
		_, err = buf.WriteString(fmt.Sprintf("\"%+v\"", k))
		if err != nil {
			return true
		}
		err = buf.WriteByte(':')
		if err != nil {
			return true
		}
		var vs []byte
		vs, err = json.Marshal(v)
		if err != nil {
			return true
		}
		_, err = buf.Write(vs)
		if err != nil {
			return true
		}
		if i < int(self.data.Count())-1 {
			err = buf.WriteByte(',')
			if err != nil {
				return true
			}
		}
		i++
		return false
	})
	if err != nil {
		return nil, err
	}
	err = buf.WriteByte('}')
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (self *_SwissTable[K, V]) Length() uint {
	return uint(self.data.Count())
}

// Set 插入键值对
func (self *_SwissTable[K, V]) Set(k K, v V) V {
	oldV, _ := self.data.Get(k)
	self.data.Put(k, v)
	return oldV
}

// Get 获取值
func (self *_SwissTable[K, V]) Get(k K, defaultValue ...V) V {
	v, ok := self.data.Get(k)
	if !ok {
		return stlslices.Last(defaultValue)
	}
	return v
}

// Contain 是否包含键
func (self *_SwissTable[K, V]) Contain(k K) bool {
	return self.data.Has(k)
}

// Remove 移除键值对
func (self *_SwissTable[K, V]) Remove(k K, defaultValue ...V) V {
	v, ok := self.data.Get(k)
	if !ok {
		return stlslices.Last(defaultValue)
	}
	self.data.Delete(k)
	return v
}

// Clear 清空
func (self *_SwissTable[K, V]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_SwissTable[K, V]) Empty() bool {
	return self.data.Count() == 0
}

// Keys 获取所有键
func (self *_SwissTable[K, V]) Keys() []K {
	ks := make([]K, self.Length())
	var i int
	self.data.Iter(func(k K, _ V) bool {
		ks[i] = k
		i++
		return false
	})
	return ks
}

// Values 获取所有值
func (self *_SwissTable[K, V]) Values() []V {
	vs := make([]V, self.Length())
	var i int
	self.data.Iter(func(_ K, v V) bool {
		vs[i] = v
		i++
		return false
	})
	return vs
}

// KeyValues 获取所有键值对
func (self *_SwissTable[K, V]) KeyValues() []tuple.Tuple2[K, V] {
	kvs := make([]tuple.Tuple2[K, V], self.Length())
	var i int
	self.data.Iter(func(k K, v V) bool {
		kvs[i] = tuple.Pack2(k, v)
		i++
		return false
	})
	return kvs
}

func (self *_SwissTable[K, V]) String() string {
	var buf strings.Builder
	var i int
	buf.WriteString("HashMap{")
	self.data.Iter(func(k K, v V) bool {
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < int(self.data.Count())-1 {
			buf.WriteString(", ")
		}
		i++
		return false
	})
	buf.WriteByte('}')
	return buf.String()
}
