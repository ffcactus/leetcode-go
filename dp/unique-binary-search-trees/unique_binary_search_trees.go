package unique_binary_search_trees

// https://leetcode-cn.com/problems/unique-binary-search-trees/

// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

// 从数组中任意取一点，把该点左边的数做成二叉搜索树，右边的也做成二叉搜索树，那么左边和右边的笛卡尔乘积就等于该root点所有的情况。
// 把每个点都做一次root，那么方法之和就是本题的解。
// 左边的情形、右边的情形有等价于原问题。那么可用动态规划来求解。
// 极端情况，如果左边（或右边）没有点或只有一个点时，我们可以直接得到答案。

const (
	Unknown = -1
)

func numTrees(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = Unknown
	}
	memo[0] = 1
	memo[1] = 1
	return dp(n, memo)
}

func dp(n int, memo []int) int {
	if memo[n] != Unknown {
		return memo[n]
	}
	sum := 0
	for i := 1; i <= n; i++ {
		left := dp(i-1, memo)
		right := dp(n-i, memo)
		sum += left * right
	}
	memo[n] = sum
	return sum
}
