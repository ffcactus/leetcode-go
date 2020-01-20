package build_tree_from_pre_in

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree return a tree from preorder and inorder values.
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

func build(preOrder []int, preLeft, preRight int, inOrder []int, inLeft, inRight int) *TreeNode {
	if preLeft == preRight {
		return nil
	}
	root := TreeNode{
		Val: preOrder[preLeft],
	}
	fmt.Println(preLeft, preRight, inLeft, inRight, root.Val)
	// i is the index that cut the inOrder array into 2 subarray.
	i := inLeft
	for ; i < inRight; i++ {
		if inOrder[i] == root.Val {
			break
		}
	}
	// the number of nodes on the left.
	leftLen := i - inLeft
	// the number of nodes on the right.
	rightLen := inRight - i - 1
	root.Left = build(
		preOrder,
		preLeft+1,
		preLeft+leftLen+1, // so when leftLen == 0, the recursion will stop.
		inOrder,
		inLeft,
		inLeft+leftLen)
	root.Right = build(
		preOrder,
		preRight-rightLen, // so when rightLen == 0, the recursion will stop.
		preRight,
		inOrder,
		inRight-rightLen,
		inRight)
	return &root
}
