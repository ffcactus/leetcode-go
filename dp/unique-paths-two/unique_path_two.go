package unique_paths_two

// https://leetcode-cn.com/problems/unique-paths-ii/

// 在不考虑障碍物的情况下：
// 到达某点的所有路径 = 到达该点上方的所有路径 + 到达改点左边所有路径。
// 我们可以直接得到地图最上方每个点的所有路径，也可以得到地图有最左边每个点的所有路径。
// 根据以上利用动态规划求解。
// 在考虑障碍物的情况下：
// 如果左边是障碍物，则只需要计算上面的点的路径数。
// 如果上边是障碍物，则只需要计算左边的点的路径数。
// 初始化时需要注意，如果第一行或第一列中有障碍物，那么后面的格子标为0，不可达。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	return search(m-1, n-1, dp, obstacleGrid)
}

func search(m, n int, dp [][]int, obstacleGrid [][]int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if obstacleGrid[m][n] == 1 {
		return 0
	}
	if dp[m][n] == 0 {
		// 如果左边有障碍物
		if n > 0 && m > 0 && obstacleGrid[m][n-1] == 1 {
			dp[m][n] = search(m-1, n, dp, obstacleGrid)
			return dp[m][n]
		}
		if n > 0 && m > 0 && obstacleGrid[m-1][n] == 1 {
			dp[m][n] = search(m, n-1, dp, obstacleGrid)
			return dp[m][n]
		}
		dp[m][n] = search(m-1, n, dp, obstacleGrid) + search(m, n-1, dp, obstacleGrid)
		return dp[m][n]
	}
	return dp[m][n]
}
