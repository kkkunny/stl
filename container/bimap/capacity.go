package bimap

func (self BiMap[K, V]) Capacity() uint {
	return self.keys.Capacity()
}
