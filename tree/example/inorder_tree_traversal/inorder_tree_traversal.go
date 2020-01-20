package inorder_tree_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	traversal(root, &ret)
	return ret
}

func traversal(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traversal(root.Left, ret)
	}
	*ret = append(*ret, root.Val)
	if root.Right != nil {
		traversal(root.Right, ret)
	}
}
