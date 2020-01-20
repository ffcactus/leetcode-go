package dfs

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsPost(root *TreeNode) {
	var (
		stack, order []int
	)
	if root == nil {
		return
	}
	dfsPostWithStack(root, &stack, &order)
	fmt.Printf("Order = %v", order)
}

func dfsPostWithStack(root *TreeNode, stack *[]int, order *[]int) {
	if root == nil {
		return
	}
	*stack = append([]int{root.Val}, *stack...)
	fmt.Printf("node = %d, stack = %v\n", root.Val, *stack)
	dfsPostWithStack(root.Left, stack, order)
	dfsPostWithStack(root.Right, stack, order)
	*order = append(*order, root.Val)
	*stack = (*stack)[1:]
}
