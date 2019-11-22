package house_robber_ii

import "fmt"

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*
为了使得问题简化，我们把圆环拉成直线，只是现在直线长度扩大一倍，即两个原数组相加，这样就不会越界。
例如原数组是 1， 2， 3， 4， 5。
现在变为 1, 2, 3, 4, 5, 1, 2, 3, 4, 5.
当我们考察从原数组中取一个数并考察有这个数时的最大值时（比如说取1），那么这个问题就变为在新数组中考察 该数字 + 随后若干位连续数字的最大值（比如说3,4）
把原数组中每个数字都考察一遍，找到最大值就是该问题的解。

 */

type key struct {
	from int
	to   int
}

func rob(nums []int) int {
	preLen := len(nums)
	nums = append(nums, nums...)

	if preLen == 1 {
		return nums[0]
	}
	if preLen == 2 {
		if nums[0] >nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	memo := make(map[key]int)
	for i := 0; i < len(nums); i++ {
		k := key{
			from: i,
			to:   i,
		}
		memo[k] = nums[i]
	}
	max := 0;
	for i:=0; i < preLen; i++ {
		right := dp(nums, i + 2, (i +2) + (preLen - 4), memo)
		m := nums[i] + right
		if m >max {
			max = m
		}
	}
	return max
}

func dp(nums []int, from, to int, memo map[key]int) int {
	fmt.Printf("checking from %d, %d = %v\n", from, to, nums[from:to + 1])
	k := key{
		from: from,
		to:   to,
	}
	if v, exist := memo[k]; exist {
		// fmt.Printf("from memo = %d\n", v)
		return v
	}
	if from+1 == to {
		if nums[from] > nums[to] {
			memo[k] = nums[from]
			// fmt.Printf("by calc = %d\n", nums[from])
			return nums[from]
		} else {
			memo[k] = nums[to]
			// fmt.Printf("by calc = %d\n", nums[to])
			return nums[to]
		}
	}
	max := 0
	for i := from; i < to; i++ {
		// fmt.Printf("check i = %d\n", i)
		// left
		left := 0
		if i > from+1 {
			left = dp(nums, from, i-2, memo)
		}
		right := 0
		if i < to-1 {
			right = dp(nums, i+2, to, memo)
		}
		m := left + nums[i] + right

		if m > max {
			max = m
		}
	}
	k = key{
		from: from,
		to:   to,
	}
	memo[k] = max
	// fmt.Printf("by calc = %d\n", max)
	return max
}