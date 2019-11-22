package swapnum

const (
	// NoNum means not exist.
	NoNum = -1
)

// 1,  6, 4, 10, 2, 5
// -1, 1, 1, 4,  1, 2

func leftNum(nums []int) []int {
	var ret = make([]int, len(nums))
	t := 0
	for i := range nums {
		if i == 0 {
			ret[t] = NoNum
			t++
			continue
		}
		if ret[i-1] == NoNum {
			if nums[i-1] < nums[i] {
				ret[t] = nums[i-1]
				t++
			} else {
				ret[t] = NoNum
				t++
			}
		} else {
			if nums[i-1] < nums[i] {
				ret[t] = nums[i-1]
				t++
			} else {
				for j := i - 1; j >= 0; j-- {
					if ret[j] < nums[i] {
						ret[t] = ret[j]
						t++
						break
					}
				}
			}
		}
	}
	return ret
}
