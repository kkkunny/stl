package _heap

import stlbasic "github.com/kkkunny/stl/basic"

type Heap[T any] struct {
	Max  bool
	Data []T
}

func (h Heap[T]) Less(i, j int) bool {
	if h.Max {
		return stlbasic.Order(h.Data[i], h.Data[j]) > 0
	} else {
		return stlbasic.Order(h.Data[i], h.Data[j]) < 0
	}
}

func (h *Heap[T]) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h Heap[T]) Len() int {
	return len(h.Data)
}

func (h *Heap[T]) Pop() (v T) {
	h.Data, v = h.Data[:h.Len()-1], h.Data[h.Len()-1]
	return
}

func (h *Heap[T]) Push(v T) {
	h.Data = append(h.Data, v)
}
