package advantagecount

import (
	"sort"
)

type record struct {
	index int
	num   int
}

func advantageCount(A []int, B []int) []int {
	arr := make([]record, 0)
	for i, v := range B {
		arr = append(arr, record{
			index: i,
			num:   v,
		})
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].num < arr[j].num
	})
	r1 := len(A) - 1
	r2 := len(B) - 1
	l1 := 0
	l2 := 0
	for l1 <= r1 {
		if A[r1] > arr[r2].num {
			B[arr[r2].index] = A[r1]
			r1--
			r2--
		} else if A[r1] < arr[r2].num {
			B[arr[r2].index] = A[l1]
			r2--
			l1++
		} else {
			if A[l1] > arr[l2].num {
				B[arr[l2].index] = A[l1]
				l1++
				l2++
			} else {
				B[arr[r2].index] = A[l1]
				l1++
				r2--
			}
		}
	}
	return B
}
