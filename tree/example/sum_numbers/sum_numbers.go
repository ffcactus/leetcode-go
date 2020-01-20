package sum_numbers

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	stack := make([]int, 0)
	dfs(root, &stack, &sum)
	return sum
}

func dfs(root *TreeNode, path *[]int, sum *int) {
	// printPath(path)
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		*sum += sumPath(path)
	}
	dfs(root.Left, path, sum)
	dfs(root.Right, path, sum)
	size := len(*path) - 1
	*path = (*path)[:size]

}

func sumPath(path *[]int) int {
	s := ""
	for _, v := range *path {
		s += fmt.Sprintf("%d", v)
	}
	v, _ := strconv.Atoi(s)
	return v
}

func printPath(path *[]int) {
	s := ""
	for _, v := range *path {
		s += fmt.Sprintf("%d ", v)
	}
	fmt.Println(s)
}
