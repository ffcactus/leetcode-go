package sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if size == 0 {
		return make([]int, 0)
	}
	left := 0
	right := k
	ret := make([]int, 0)
	for right <= size {
		max := nums[left]
		for _, v := range nums[left:right] {
			if v > max {
				max = v
			}
		}
		ret = append(ret, max)
		right++
		left++
	}
	return ret
}
