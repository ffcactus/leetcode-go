package flatten

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	recursive(root)
}

func recursive(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		r := root.Right
		root.Right = root.Left
		root.Left = nil
		var rl *TreeNode
		for rl = root.Right; rl.Right != nil; rl = rl.Right {
		}
		rl.Right = r
	}
	if root.Right != nil {
		recursive(root.Right)
	}
}
