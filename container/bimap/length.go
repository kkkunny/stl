package bimap

func (self BiMap[K, V]) Length() uint {
	return self.keys.Length()
}
