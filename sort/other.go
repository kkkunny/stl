package sort

import (
	"golang.org/x/exp/constraints"
)

// 归并排序 O(NlogN)
func MergeSort[T constraints.Ordered](l []T, reverse bool) {
	if len(l) == 1 {
		return
	}
	mid := len(l) / 2
	left, right := make([]T, mid), make([]T, len(l)-mid)
	copy(left, l[:mid])
	copy(right, l[mid:])
	MergeSort(left, reverse)
	MergeSort(right, reverse)
	var li, ri int
	for i := 0; i < len(l); i++ {
		if li == len(left) {
			l[i] = right[ri]
			ri++
		} else if ri == len(right) {
			l[i] = left[li]
			li++
		} else if (!reverse && left[li] < right[ri]) ||
			(reverse && left[li] > right[ri]) {
			l[i] = left[li]
			li++
		} else {
			l[i] = right[ri]
			ri++
		}
	}
}
