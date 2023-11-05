package linkedhashmap

func (self LinkedHashMap[K, V]) Capacity() uint {
	self.init()
	return self.HashMap.Capacity()
}
