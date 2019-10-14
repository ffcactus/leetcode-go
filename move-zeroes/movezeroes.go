package move_zeroes

func moveZeroes(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}

	start := 0
	end := len(nums) - 1

	for start < end {
		if nums[start] == 0 {
			for i := start; i < end; i++ {
				nums[i] = nums[i+1]
			}
			nums[end] = 0
			end--
		} else {
			start++
		}
	}
}
