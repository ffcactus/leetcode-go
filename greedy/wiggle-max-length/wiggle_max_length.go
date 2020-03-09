package length

const (
	up    = 1
	down  = 2
	equal = 0
)

func wiggleMaxLength(nums []int) int {
	direction := equal
	N := len(nums)
	if N < 2 {
		return N
	}
	if N == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return 2
	}
	count := 0
	if nums[0] < nums[1] {
		direction = up
		count = 2
	} else if nums[0] > nums[1] {
		direction = down
		count = 2
	} else {
		direction = equal
		count = 1
	}
	for i := 2; i < N; i++ {
		if direction == up {
			if nums[i] >= nums[i-1] {
				continue
			} else {
				count++
				direction = down
			}
		} else if direction == down {
			if nums[i] <= nums[i-1] {
				continue
			} else {
				count++
				direction = up
			}
		} else {
			if nums[i] > nums[i-1] {
				count++
				direction = up
			} else if nums[i] < nums[i-1] {
				count++
				direction = down
			}
		}
	}
	return count
}
