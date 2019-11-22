package question2

const (
	default2 = 2
)

func findShortestDuration(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == default2 {
		if nums[0] == nums[1] {
			return default2
		} else {
			return 1
		}
	}
	st := make(map[int]int)
	for _, v := range nums {
		st[v]++
	}
	// 记录最多的噪声数。
	max := 0
	for _, v := range st {
		if v > max {
			max = v
		}
	}
	// 记录哪些频段满足最大的噪声数。
	var ks []int
	for k, v := range st {
		if v == max {
			ks = append(ks, k)
		}
	}
	// fmt.Println("max", max, "ks", ks)
	left := 0
	right := 1
	min := len(nums)
	for true {
		// fmt.Println("left", left, "right", right)
		if match(nums[left:right+1], ks, max, min) {
			if len(nums[left:right+1]) < min {
				min = len(nums[left : right+1])
			}
			left++
		} else {
			right++
			if right == len(nums) {
				right--
				left++
			}
		}
		if left == right && right == len(nums)-1 {
			break
		}
	}
	return min
}

func match(nums []int, ks []int, max, min int) bool {
	if len(nums) > min {
		return false
	}
	// 如果有K个MAX则满足
	for _, k := range ks {
		count := 0
		for _, n := range nums {
			if n == k {
				count++
			}
		}
		if count >= max {
			// fmt.Println("match()", "nums", nums, "true")
			return true
		}
	}
	// fmt.Println("match()", "nums", nums, "false")
	return false
}
