package longest_consecutive_sequence

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	sort.Ints(nums)
	max := 1
	for i := range nums {
		tmpMax := 1
		for j := i; j < len(nums)-1; j++ {
			if nums[j]+1 == nums[j+1] {
				tmpMax++
			} else if nums[j] == nums[j+1] {
				continue
			} else {
				break
			}
		}
		if tmpMax > max {
			max = tmpMax
		}
	}
	return max
}
