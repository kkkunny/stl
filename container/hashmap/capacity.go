package hashmap

func (self HashMap[K, V]) Capacity() uint {
	self.init()
	return uint(cap(*self.buckets))
}
