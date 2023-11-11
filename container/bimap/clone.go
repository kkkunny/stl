package bimap

func (self BiMap[K, V]) Clone() BiMap[K, V] {
	return BiMap[K, V]{
		keys:   self.keys.Clone(),
		values: self.values.Clone(),
	}
}
