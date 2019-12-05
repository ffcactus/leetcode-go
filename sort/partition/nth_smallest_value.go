package partition

// nthSmallest select the n-th smallest value from the array a.
// for example, the 3-th smallest value from 5, 4, 3, 2, 1, 0 is 3.
func nthSmallest(a []int, n int) int {
	var (
		low = 0
		high = len(a) - 1
	)

	for high > low {
		j := partition(a, low, high)
		if j == n {
			return a[j]
		} else if j > n {
			high = j - 1;
		} else if j < n {
			low = j + 1
		}
	}
	return a[n]
}