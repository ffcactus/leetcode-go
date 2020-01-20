package is_valid_bst

import (
	"math"

	"leetcode-go/tree/example/flatten"
)

func isValidBST(root *flatten.TreeNode) bool {
	return check(root, math.MinInt64, math.MinInt64)
}

func check(root *flatten.TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if low != math.MinInt64 && root.Val <= low {
		return false
	}
	if high != math.MinInt64 && root.Val >= high {
		return false
	}
	if !check(root.Left, low, root.Val) {
		return false
	}
	if !check(root.Right, root.Val, high) {
		return false
	}
	return true
}
