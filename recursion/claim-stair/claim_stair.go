package stair

func climbStairs(n int) int {
	memo := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	memo[1] = 1
	memo[2] = 2
	return dp(n, memo)
}

func dp(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = dp(n-1, memo) + dp(n-2, memo)
	return memo[n]
}
