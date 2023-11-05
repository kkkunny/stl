package hashmap

func (self HashMap[K, V]) Length() uint {
	self.init()
	return self.length
}
