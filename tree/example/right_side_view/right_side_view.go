package right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	nodes := levelOrder(root)
	ret := make([]int, 0)
	for _, v := range nodes {
		ret = append(ret, v[len(v)-1])
	}
	return ret
}

// levelOrder returns the an array for each of the level.
func levelOrder(root *TreeNode) [][]int {
	var (
		values *[]int
	)
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	next := []*TreeNode{root}
	for {
		if len(next) != 0 {
			next, values = children(next)
			ret = append(ret, *values)
		} else {
			break
		}
	}
	return ret
}

func children(roots []*TreeNode) ([]*TreeNode, *[]int) {
	collection := make([]*TreeNode, 0)
	values := make([]int, 0)
	for _, n := range roots {
		if n.Left != nil {
			collection = append(collection, n.Left)
		}
		if n.Right != nil {
			collection = append(collection, n.Right)
		}
		values = append(values, n.Val)
	}
	return collection, &values
}
