package unique_paths_three

/*
在二维网格 grid 上，有 4 种类型的方格：

1 表示起始方格。且只有一个起始方格。
2 表示结束方格，且只有一个结束方格。
0 表示我们可以走过的空方格。
-1 表示我们无法跨越的障碍。
返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目，每一个无障碍方格都要通过一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 显然，目标点的路径数等于上下左右4个点的路径数之和。
// 题目要求每个点必须通过一次，且只能通过一次，因此既需要记录路径。
func uniquePathsIII(grid [][]int) int {
	return 0
}

func search(x0, y0, x1, y1, maxX, maxY int, grid [][]int) int {

	sum := 0
	// 上
	if y1 > 0 {
		if grid[y1-1][x1] == 0 {
			newGrid := grid
			newGrid[y1][x1] = -1
			sum += search(x0, y0, x1, y1-1, maxX, maxY, newGrid)
		}
	}
	// 下
	if y1 < maxY {
		if grid[y1+1][x1] == 0 {
			newGrid := grid
			newGrid[y1][x1] = -1
			sum += search(x0, y0, x1, y1+1, maxX, maxY, newGrid)
		}
	}
	// 左
	if x1 > 0 {
		if grid[y1][x1-1] == 0 {
			newGrid := grid
			newGrid[y1][x1] = -1
			sum += search(x0, y0, x1-1, y1, maxX, maxY, newGrid)
		}
	}
	// 右
	if x1 < maxX {
		if grid[y1][x1+1] == 0 {
			newGrid := grid
			newGrid[y1][x1] = -1
			sum += search(x0, y0, x1+1, y1, maxX, maxY, newGrid)
		}
	}
	return sum
}
