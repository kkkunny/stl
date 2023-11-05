package hashmap

func (self HashMap[K, V]) Clone() HashMap[K, V] {
	newData := make([]bucket[K, V], len(*self.buckets), cap(*self.buckets))
	copy(newData, *self.buckets)
	return HashMap[K, V]{
		buckets: &newData,
		length:  self.length,
	}
}
