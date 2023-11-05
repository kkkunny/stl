package hashset

func (self HashSet[T]) Length() uint {
	return self.data.Length()
}
