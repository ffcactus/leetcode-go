package wiggle_sort

import "sort"

func wiggleSort(nums []int) {
	var (
		left, right []int
	)
	sort.Ints(nums)
	if len(nums)%2 == 1 {
		left = nums[:len(nums)/2+1]
		right = nums[len(nums)/2+1:]
		sort.Slice(left, func(i, j int) bool {
			return left[i] > left[j]
		})
		sort.Slice(right, func(i, j int) bool {
			return right[i] > right[j]
		})
		ret := make([]int, len(nums))
		for i := range right {
			ret[i*2] = left[i]
			ret[i*2+1] = right[i]
		}
		ret[2*len(right)] = left[len(left)-1]
		for i := range ret {
			nums[i] = ret[i]
		}
	} else {
		left = nums[:len(nums)/2]
		right = nums[len(nums)/2:]
		sort.Slice(left, func(i, j int) bool {
			return left[i] > left[j]
		})
		sort.Slice(right, func(i, j int) bool {
			return right[i] > right[j]
		})
		ret := make([]int, len(nums))
		for i := range right {
			ret[i*2] = left[i]
			ret[i*2+1] = right[i]
		}
		for i := range ret {
			nums[i] = ret[i]
		}
	}
}
