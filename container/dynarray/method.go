package dynarray

// 初始化
func (self *DynArray[T]) init() {
	if self.data != nil {
		return
	}
	self.data = new([]T)
}

// Get 获取元素
func (self DynArray[T]) Get(i uint) T {
	self.init()
	return (*self.data)[i]
}

// Set 设置元素
func (self *DynArray[T]) Set(i uint, v T) T {
	self.init()
	pv := (*self.data)[i]
	(*self.data)[i] = v
	return pv
}

// PushBack 插入新元素到末尾
func (self *DynArray[T]) PushBack(v T, vs ...T) {
	self.init()
	*self.data = append(*self.data, v)
	*self.data = append(*self.data, vs...)
}

// PushFront 插入新元素到开头
func (self *DynArray[T]) PushFront(v T, vs ...T) {
	self.init()
	*self.data = append(append([]T{v}, vs...), *self.data...)
}

// Insert 插入新元素
func (self *DynArray[T]) Insert(i uint, v T, vs ...T) {
	self.init()
	*self.data = append(*self.data, v)
	*self.data = append(*self.data, vs...)
	copy((*self.data)[i+1+uint(len(vs)):], (*self.data)[i:])
	(*self.data)[i] = v
	for vi, vv := range vs {
		(*self.data)[i+uint(vi)+1] = vv
	}
}

// Remove 移除元素
func (self *DynArray[T]) Remove(i uint) T {
	self.init()
	v := (*self.data)[i]
	copy((*self.data)[i:], (*self.data)[i+1:])
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

// PopBack 移除末尾元素
func (self *DynArray[T]) PopBack() T {
	self.init()
	v := (*self.data)[len(*self.data)-1]
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

// PopFront 移除开头元素
func (self *DynArray[T]) PopFront() T {
	self.init()
	v := (*self.data)[0]
	*self.data = (*self.data)[1:]
	return v
}

// Back 获取末尾元素
func (self DynArray[T]) Back() T {
	self.init()
	return (*self.data)[len(*self.data)-1]
}

// Front 获取开头元素
func (self DynArray[T]) Front() T {
	self.init()
	return (*self.data)[0]
}

// Clear 清空
func (self *DynArray[T]) Clear() {
	self.init()
	*self.data = make([]T, 0)
}

// Empty 是否为空
func (self DynArray[T]) Empty() bool {
	self.init()
	return self.Length() == 0
}
