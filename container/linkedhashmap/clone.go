package linkedhashmap

func (self LinkedHashMap[K, V]) Clone() LinkedHashMap[K, V] {
	self.init()
	hm := NewLinkedHashMapWithCapacity[K, V](self.Capacity())
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		hm.Set(cursor.Value.First, cursor.Value.Second)
	}
	return hm
}
