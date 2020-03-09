package smallrangeii

import (
	"sort"
)

func smallestRangeII(A []int, K int) int {
	N := len(A)
	sort.Ints(A)
	ans := A[N-1] - A[0]
	for i := 0; i < N-1; i++ {
		a := A[i]
		b := A[i+1]
		hight := max(A[N-1]-K, a+K)
		low := min(A[0]+K, b-K)
		ans = min(ans, hight-low)
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
