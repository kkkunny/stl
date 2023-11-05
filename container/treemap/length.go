package treemap

func (self TreeMap[K, V]) Length() uint {
	self.init()
	return self.tree.Len()
}
