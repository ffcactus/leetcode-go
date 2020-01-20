package build_tree_from_in_post

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree build tree from in-order and post-order array.
func buildTree(inOrder []int, postOrder []int) *TreeNode {
	return build(inOrder, 0, len(inOrder)-1, postOrder, 0, len(postOrder)-1)
}

func build(inOrder []int, inFrom, inTo int, postOrder []int, postFrom, postTo int) *TreeNode {

	if postFrom > postTo {
		return nil
	}

	root := &TreeNode{
		Val: postOrder[postTo],
	}
	rootIndex := inFrom
	for ; rootIndex <= inTo; rootIndex++ {
		if inOrder[rootIndex] == root.Val {
			break
		}
	}
	leftCount := rootIndex - inFrom
	// rightCount := inTo - rootIndex
	root.Left = build(
		inOrder,
		inFrom,
		inFrom+leftCount-1,
		postOrder,
		postFrom,
		postFrom+leftCount-1)
	root.Right = build(
		inOrder,
		rootIndex+1,
		inTo,
		postOrder,
		postFrom+leftCount,
		postTo-1)
	return root
}
