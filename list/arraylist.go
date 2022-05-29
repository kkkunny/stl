package list

import (
	"encoding/json"
	"fmt"
)

// 动态数组
type ArrayList[T any] struct {
	data []T
}

// 新建动态数组
func NewArrayList[T any](length, capacity int) *ArrayList[T] {
	return &ArrayList[T]{
		data: make([]T, length, capacity),
	}
}

// 新建动态数组带初始值
func NewArrayListWithInitial[T any](e ...T) *ArrayList[T] {
	return &ArrayList[T]{
		data: e,
	}
}

// 根据切片新建动态数组
func NewArrayListFromSlice[T any](slice []T) *ArrayList[T] {
	return &ArrayList[T]{
		data: slice,
	}
}

// 转成字符串 O(N)
func (self *ArrayList[T]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// json序列化
func (self *ArrayList[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.data)
}

// json反序列化
func (self *ArrayList[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &self.data)
}

// 获取长度 O(1)
func (self *ArrayList[T]) Length() int {
	return len(self.data)
}

// 获取容量 O(1)
func (self *ArrayList[T]) Capacity() int {
	return cap(self.data)
}

// 是否为空 O(1)
func (self *ArrayList[T]) Empty() bool {
	return len(self.data) == 0
}

// 增加元素 O(1)-O(N)
func (self *ArrayList[T]) Add(e ...T) {
	self.data = append(self.data, e...)
}

// 插入元素 O(N)
func (self *ArrayList[T]) Insert(i int, e ...T) {
	self.data = append(self.data[:i], append(e, self.data[i:]...)...)
}

// 移除元素 O(N)
func (self *ArrayList[T]) Remove(i int) T {
	elem := self.data[i]
	self.data = append(self.data[:i], self.data[i+1:]...)
	return elem
}

// 获取元素 O(1)
func (self *ArrayList[T]) Get(i int) T {
	return self.data[i]
}

// 删除头元素 O(1)
func (self *ArrayList[T]) PopFront() T {
	elem := self.data[0]
	self.data = self.data[1:]
	return elem
}

// 删除尾元素 O(1)
func (self *ArrayList[T]) PopBack() T {
	elem := self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]
	return elem
}

// 获取第一个元素 O(1)
func (self *ArrayList[T]) First() T {
	return self.data[0]
}

// 获取最后一个元素 O(1)
func (self *ArrayList[T]) Last() T {
	return self.data[len(self.data)-1]
}

// 设置元素 O(1)
func (self *ArrayList[T]) Set(i int, e T) T {
	elem := self.data[i]
	self.data[i] = e
	return elem
}

// 清空 O(1)
func (self *ArrayList[T]) Clear() {
	if self.Empty() {
		return
	}
	self.data = nil
}

// 克隆 O(N)
func (self *ArrayList[T]) Clone() *ArrayList[T] {
	data := make([]T, len(self.data))
	copy(data, self.data)
	return &ArrayList[T]{
		data: data,
	}
}

// 过滤 O(N)
func (self *ArrayList[T]) Filter(f func(i int, v T) bool) *ArrayList[T] {
	al := NewArrayList[T](0, 0)
	for i, v := range self.data {
		if f(i, v) {
			al.Add(v)
		}
	}
	return al
}

// 切分[b, e) O(N)
func (self *ArrayList[T]) Slice(b, e int) *ArrayList[T] {
	tmp := self.data[b:e]
	a := NewArrayList[T](len(tmp), len(tmp))
	copy(a.data, tmp)
	return a
}

// 拼接 O(N)
func (self *ArrayList[T]) Contact(a *ArrayList[T]) {
	self.data = append(self.data, a.data...)
}

// 是否有任何元素满足条件 O(N)
func (self *ArrayList[T]) Any(f func(i int, v T) bool) bool {
	for i, v := range self.data {
		if f(i, v) {
			return true
		}
	}
	return false
}

// 是否所有元素都满足条件 O(N)
func (self *ArrayList[T]) Every(f func(i int, v T) bool) bool {
	for i, v := range self.data {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// 获取起始迭代器
func (self *ArrayList[T]) Begin() *ArrayListIterator[T] {
	return &ArrayListIterator[T]{data: self.data}
}

// 获取结束迭代器
func (self *ArrayList[T]) End() *ArrayListIterator[T] {
	return &ArrayListIterator[T]{
		data:  self.data,
		index: self.Length() - 1,
	}
}

// 迭代器
type ArrayListIterator[T any] struct {
	data  []T // 列表
	index int // 目前索引
}

// 是否存在值
func (self *ArrayListIterator[T]) HasValue() bool {
	return 0 <= self.index && self.index < len(self.data)
}

// 是否存在上一个
func (self *ArrayListIterator[T]) HasPrev() bool {
	return self.index >= 1
}

// 是否存在下一个
func (self *ArrayListIterator[T]) HasNext() bool {
	return self.index+1 < len(self.data)
}

// 上一个
func (self *ArrayListIterator[T]) Prev() {
	self.index--
}

// 下一个
func (self *ArrayListIterator[T]) Next() {
	self.index++
}

// 获取索引
func (self *ArrayListIterator[T]) Index() int {
	return self.index
}

// 获取值
func (self *ArrayListIterator[T]) Value() T {
	return self.data[self.index]
}
