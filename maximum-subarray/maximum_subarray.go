package maximumsubarray

// https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	if len == 0 {
		return 0
	}
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
