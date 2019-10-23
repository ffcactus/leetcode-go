package _sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	return impl2(nums)
}

func impl1(nums []int) [][]int {
	memo := make(map[string]bool)
	var ret [][]int
	size := len(nums)
	if size < 3 {
		return ret
	}
	sort.Ints(nums)
	for i:=0; i < size - 2; i++ {
		for j:=i + 1; j <size - 1; j++ {
			for k:=j+1; k <size; k++ {

				if nums[i] + nums[j] + nums[k] == 0 {
					s := fmt.Sprintf("%d-%d-%d", nums[i], nums[j], nums[k])
					if _, ok := memo[s]; !ok {
						memo[s] = true
						ret = append(ret, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return ret
}

func impl2(nums []int) [][]int {
	memo := make(map[string]bool)
	var ret [][]int
	size := len(nums)
	if size < 3 {
		return ret
	}
	sort.Ints(nums)
	zeroIndex := 0
	for i:=0; i < size; i++ {
		if nums[i] > 0 {
			zeroIndex = i
			break
		}
	}

	for i:=0; i < size - 2; i++ {
		for j:=i + 1; j <size - 1; j++ {
			ki := j+1
			if zeroIndex > ki {
				if nums[i] + nums[j] < 0 {
					ki = zeroIndex
				}
			}
			for k:=ki; k <size; k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					s := fmt.Sprintf("%d-%d-%d", nums[i], nums[j], nums[k])
					if _, ok := memo[s]; !ok {
						memo[s] = true
						ret = append(ret, []int{nums[i], nums[j], nums[k]})
					}
					break
				}
			}
		}
	}
	return ret
}
