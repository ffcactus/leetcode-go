package zigzag_level_order

import (
	"reflect"

	"leetcode-go/tree/example/flatten"
)

// zigzagLevelOrder return an array for each level of tree root.
// The first level from left to right.
// The second level from right to left.
// And so on.
func zigzagLevelOrder(root *flatten.TreeNode) [][]int {
	var (
		values *[]int
	)
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	next := []*flatten.TreeNode{root}
	fromLeft := true
	for {
		if len(next) != 0 {
			next, values = zigzag(next, fromLeft)
			fromLeft = !fromLeft
			ret = append(ret, *values)
		} else {
			break
		}
	}
	return ret
}

func zigzag(roots []*flatten.TreeNode, fromLeft bool) ([]*flatten.TreeNode, *[]int) {
	collection := make([]*flatten.TreeNode, 0)
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

	if !fromLeft {
		reverseAny(values)
	}
	return collection, &values
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
