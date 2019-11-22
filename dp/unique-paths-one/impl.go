package unique_paths_one

// https://leetcode-cn.com/problems/unique-paths/
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？

// 到达某点的所有路径 = 到达该点上方的所有路径 + 到达改点左边所有路径。
// 我们可以直接得到地图最上方每个点的所有路径，也可以得到地图有最左边每个点的所有路径。
// 根据以上利用动态规划求解。
func uniquePaths(m int, n int) int {
	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	return search(m-1, n-1, dp)
}

func search(m, n int, dp [][]int) int {
	if dp[m][n] == 0 {
		dp[m][n] = search(m-1, n, dp) + search(m, n-1, dp)
		return dp[m][n]
	}
	return dp[m][n]
}
