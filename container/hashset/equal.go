package hashset

func (self HashSet[T]) Equal(dst HashSet[T]) bool {
	return self.data.Equal(dst.data)
}
