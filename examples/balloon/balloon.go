package balloon

// https://leetcode-cn.com/problems/burst-balloons/

// 如何是的最终结果最大化？
// 假设i,j是位置，原问题是求nums[i:j], 当i=0; j=n时的最大值。
// 从i,j中找到一个k，原问题变为 i->k, k, k->j 三个值的和的最大值。其中i->k 和 k-> 分别代表左边全部戳破和右边全部戳破得到的最大值。
// 左右全部戳破后，k的值为num[i] * num[k] * num[j]
// 动态转移方程：dp[i][j] = math.Max(dp[i][j],   dp[i][k] + dp[k][j] + num[i] * num[k] * num[j])
func maxCoins(nums []int) int {

	// 构建 dp = [len(nums) + 2][len(nums) + 2]int
	// +2 意味着左右各加一个，即nums[-1] = nums[len(nums)] = 1, 方便计算。
	dp := make([][]int, 0)
	for i := 0; i < len(nums)+2; i++ {
		dp = append(dp, make([]int, len(nums)+2))
	}

	newNums := make([]int, len(nums)+2)
	for i := 1; i < len(newNums)-1; i++ {
		newNums[i] = nums[i-1]
	}
	newNums[0] = 1
	newNums[len(newNums)-1] = 1
	// 从2开始，保证最少3个
	for j := 2; j < len(newNums); j++ {
		// 遍历所有的可能性，0-2...0-n,1-3...1-n,...
		for i := 0; i < len(newNums)-j; i++ {
			for k := i + 1; k < i+j; k++ {
				a := dp[i][i+j]
				b := dp[i][k] + dp[k][i+j] + newNums[i]*newNums[k]*newNums[i+j]
				if a > b {
					dp[i][i+j] = a
				} else {
					dp[i][i+j] = b
				}
			}
		}
	}
	return dp[0][len(newNums)-1]
}
