package contains_nearby_almost_duplicate

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := range nums {
		for j := max(i-k, 0); j < i; j++ {
			fmt.Println(nums[i], nums[j])
			if abs(nums[i], nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
