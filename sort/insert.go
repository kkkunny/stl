package sort

import (
	"golang.org/x/exp/constraints"
)

// 插入排序 O(N²)
func InsertSort[T constraints.Ordered](l []T, reverse bool) {
	for i := 1; i < len(l); i++ {
		tmp := l[i]
		j := i - 1
		for ; j >= 0 && ((!reverse && l[j] > tmp) || (reverse && l[j] < tmp)); j-- {
			l[j+1] = l[j]
		}
		l[j+1] = tmp
	}
}

// 二分插入排序 O(N²)
func BinaryInsertSort[T constraints.Ordered](l []T, reverse bool) {
	for i := 1; i < len(l); i++ {
		tmp := l[i]
		left, right := 0, i-1
		for left <= right {
			mid := left + (right-left)/2
			if (!reverse && tmp <= l[mid]) ||
				(reverse && tmp > l[mid]) {
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
func ShellSort[T constraints.Ordered](l []T, reverse bool) {
	gap := len(l) / 2
	for gap > 0 {
		for i := gap; i < len(l); i++ {
			tmp := l[i]
			j := i - gap
			for ; j >= 0 && ((!reverse && l[j] > tmp) || (reverse && l[j] < tmp)); j -= gap {
				l[j+gap] = l[j]
			}
			l[j+gap] = tmp
		}
		gap /= 2
	}
}
