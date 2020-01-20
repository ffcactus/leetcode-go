package path_sum2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var (
		paths [][]int
	)
	dfsPost(root, sum, &paths)
	return paths
}

func dfsPost(root *TreeNode, sum int, paths *[][]int) {
	var (
		stack, order []int
	)
	if root == nil {
		return
	}
	dfsPostWithStack(root, &stack, &order, sum, paths)
}

func dfsPostWithStack(root *TreeNode, stack *[]int, order *[]int, sum int, paths *[][]int) {
	if root == nil {
		return
	}
	*stack = append([]int{root.Val}, *stack...)
	// fmt.Printf("node = %d, stack = %v\n", root.Val, *stack)
	if root.Left == nil && root.Right == nil {
		s := 0
		var path []int
		for _, v := range *stack {
			path = append([]int{v}, path...)
			s += v
		}
		if s == sum {
			*paths = append(*paths, path)
		}
	}
	dfsPostWithStack(root.Left, stack, order, sum, paths)
	dfsPostWithStack(root.Right, stack, order, sum, paths)
	*order = append(*order, root.Val)
	*stack = (*stack)[1:]
}
