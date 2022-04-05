package sort

import (
	. "github.com/kkkunny/stl/types"
)

// 冒泡排序 O(N²)
func BubbleSort[T Comparator[T]](l []T, reverse bool) {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if (!reverse && l[i].Compare(l[j]) > 0) ||
				(reverse && l[i].Compare(l[j]) < 0) {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
}

// 快速排序 O(NlogN)
func QuickSort[T Comparator[T]](l []T, reverse bool) {
	switch len(l) {
	case 0, 1:
	case 2:
		if (!reverse && l[0].Compare(l[1]) > 0) ||
			(reverse && l[0].Compare(l[1]) < 0) {
			l[0], l[1] = l[1], l[0]
		}
	default:
		mid := l[len(l)/2]
		var midCount int
		var small, big []T
		for i := 0; i < len(l); i++ {
			diff := l[i].Compare(mid)
			if diff == 0 {
				midCount++
			} else if diff > 0 {
				big = append(big, l[i])
			} else {
				small = append(small, l[i])
			}
		}
		if midCount == len(l) {
			return
		}
		QuickSort(small, reverse)
		QuickSort(big, reverse)
		if !reverse {
			copy(l, small)
			for i := 0; i < midCount; i++ {
				l[len(small)+i] = mid
			}
			copy(l[len(small)+midCount:], big)
		} else {
			copy(l, big)
			for i := 0; i < midCount; i++ {
				l[len(big)+i] = mid
			}
			copy(l[len(big)+midCount:], small)
		}
	}
}
