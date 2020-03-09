package jump

const (
	can     = 1
	cannot  = 2
	unknown = 0
)

func canJump(nums []int) bool {
	memo := make([]int, len(nums))
	return canJumpFrom(0, nums, memo)
}

func canJump0(nums []int) bool {
	lastPos := len(nums) - 1 // lastPos is the most right point that can jump to the end.
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos { // if you can jump to the lastPos can you can jump to the end.
			lastPos = i
		}
	}
	return lastPos == 0
}

func canJumpFrom(i int, nums []int, memo []int) bool {
	if memo[i] == can {
		return true
	}
	if memo[i] == cannot {
		return false
	}
	N := len(nums)
	if i == N-1 {
		memo[i] = can
		return true
	}
	maxJump := min(i+nums[i], N-1)
	for j := maxJump; j > i; j-- {
		if canJumpFrom(j, nums, memo) {
			memo[i] = can
			return true
		}
	}
	memo[i] = cannot
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
