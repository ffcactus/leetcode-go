package house_robber

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
在数组中任取一个数，那么在取这个数的情况下的最大值是：隔一位的左边所有数的最大值+这个数+隔一位右边所有数的最大值。
而隔一位左边，右边所有数的最大值又等于原问题。那么遍历原数组的每一位数，就能找出所有情况下的最大值。
迭代的出口条件可以设置成当数组只有1位，2位数时。
这种算法在数组较大是容易超时，引入memo记录2个数组下标from,to之间的最大值。
*/
type key struct {
	from int
	to   int
}

func rob(nums []int) int {
	memo := make(map[key]int)
	for i := 0; i < len(nums); i++ {
		k := key{
			from: i,
			to:   i,
		}
		memo[k] = nums[i]
	}
	return dp(nums, 0, len(nums)-1, memo)
}

func dp(nums []int, from, to int, memo map[key]int) int {
	// fmt.Printf("from %d, to %d\n", from, to)
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
