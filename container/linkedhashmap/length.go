package linkedhashmap

func (self LinkedHashMap[K, V]) Length() uint {
	self.init()
	return self.HashMap.Length()
}
