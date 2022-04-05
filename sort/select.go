package sort

import (
	"github.com/kkkunny/stl/heap"
	. "github.com/kkkunny/stl/types"
)

// 选择排序 O(N²)
func SelectSort[T Comparator[T]](l []T, reverse bool) {
	for i := 0; i < len(l); i++ {
		swap := i
		for j := i + 1; j < len(l); j++ {
			if (!reverse && l[j].Compare(l[swap]) < 0) ||
				(reverse && l[j].Compare(l[swap]) > 0) {
				swap = j
			}
		}
		if swap != i {
			l[i], l[swap] = l[swap], l[i]
		}
	}
}

// 堆排序 O(NlogN)-O(N²logN)
func HeapSort[T Comparator[T]](l []T, reverse bool) {
	var h heap.Heap[T]
	if !reverse {
		h = heap.NewMinHeap[T]()
	} else {
		h = heap.NewMaxHeap[T]()
	}
	for i := 0; i < len(l); i++ {
		h.Push(l[i])
	}
	for i := 0; i < len(l); i++ {
		l[i] = h.Pop()
	}
}
