package sort

import (
	. "github.com/kkkunny/stl/types"
)

// 插入排序 O(N²)
func InsertSort[T Comparator[T]](l []T, reverse bool) {
	for i := 1; i < len(l); i++ {
		tmp := l[i]
		j := i - 1
		for ; j >= 0 && ((!reverse && l[j].Compare(tmp) > 0) || (reverse && l[j].Compare(tmp) < 0)); j-- {
			l[j+1] = l[j]
		}
		l[j+1] = tmp
	}
}

// 二分插入排序 O(N²)
func BinaryInsertSort[T Comparator[T]](l []T, reverse bool) {
	for i := 1; i < len(l); i++ {
		tmp := l[i]
		left, right := 0, i-1
		for left <= right {
			mid := left + (right-left)/2
			if (!reverse && tmp.Compare(l[mid]) <= 0) ||
				(reverse && tmp.Compare(l[mid]) >= 0) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i - 1; j > right; j-- {
			l[j+1] = l[j]
		}
		l[right+1] = tmp
	}
}

// 希尔排序 O(N**1.3) - O(N²)
func ShellSort[T Comparator[T]](l []T, reverse bool) {
	gap := len(l) / 2
	for gap > 0 {
		for i := gap; i < len(l); i++ {
			tmp := l[i]
			j := i - gap
			for ; j >= 0 && ((!reverse && l[j].Compare(tmp) > 0) || (reverse && l[j].Compare(tmp) < 0)); j -= gap {
				l[j+gap] = l[j]
			}
			l[j+gap] = tmp
		}
		gap /= 2
	}
}
