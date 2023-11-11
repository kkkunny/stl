package bimap

func (self BiMap[K, V]) Equal(dst BiMap[K, V]) bool {
	return self.keys.Equal(dst.keys)
}
