package hashset

func (self HashSet[T]) Clone() HashSet[T] {
	return HashSet[T]{data: self.data.Clone()}
}
