package list

import (
	"fmt"

	. "github.com/kkkunny/stl/types"
)

// 动态数组
type ArrayList[T any] struct {
	data []T
}

// 新建动态数组
func NewArrayList[T any](length, capacity Usize) *ArrayList[T] {
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

// 转成字符串
func (self *ArrayList[T]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// 获取长度
func (self *ArrayList[T]) Length() Usize {
	return Usize(len(self.data))
}

// 获取容量
func (self *ArrayList[T]) Capacity() Usize {
	return Usize(cap(self.data))
}

// 是否为空
func (self *ArrayList[T]) Empty() bool {
	return len(self.data) == 0
}

// 增加元素
func (self *ArrayList[T]) Add(e ...T) {
	self.data = append(self.data, e...)
}

// 插入元素
func (self *ArrayList[T]) Insert(i Usize, e ...T) {
	temp := self.data[i:]
	self.data = append(self.data[:i], e...)
	self.data = append(self.data, temp...)
}

// 移除元素
func (self *ArrayList[T]) Remove(i Usize) T {
	elem := self.data[i]
	self.data = append(self.data[:i], self.data[i+1:]...)
	return elem
}

// 获取元素
func (self *ArrayList[T]) Get(i Usize) T {
	return self.data[i]
}

// 获取第一个节点
func (self *ArrayList[T]) First() T {
	return self.data[0]
}

// 获取最后一个节点
func (self *ArrayList[T]) Last() T {
	return self.data[len(self.data)-1]
}

// 设置元素
func (self *ArrayList[T]) Set(i Usize, e T) T {
	elem := self.data[i]
	self.data[i] = e
	return elem
}

// 清空
func (self *ArrayList[T]) Clear() {
	self.data = make([]T, self.Capacity())
}

// 克隆
func (self *ArrayList[T]) Clone() *ArrayList[T] {
	data := make([]T, len(self.data), cap(self.data))
	copy(data, self.data)
	return &ArrayList[T]{
		data: data,
	}
}

// 获取起始迭代器
func (self *ArrayList[T]) Begin() *ArrayListIterator[T] {
	return &ArrayListIterator[T]{data: self}
}

// 获取结束迭代器
func (self *ArrayList[T]) End() *ArrayListIterator[T] {
	return &ArrayListIterator[T]{
		data:  self,
		index: self.Length() - 1,
	}
}

// 迭代器
type ArrayListIterator[T any] struct {
	data  *ArrayList[T] // 列表
	index Usize         // 目前索引
}

// 是否存在值
func (self *ArrayListIterator[T]) HasValue() bool {
	return 0 <= self.index && self.index < Usize(len(self.data.data))
}

// 上一个
func (self *ArrayListIterator[T]) Prev() {
	self.index--
}

// 下一个
func (self *ArrayListIterator[T]) Next() {
	self.index++
}

// 获取值
func (self *ArrayListIterator[T]) Value() T {
	return self.data.data[self.index]
}
