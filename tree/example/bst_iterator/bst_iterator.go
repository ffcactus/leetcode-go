package bst_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	index   int
	inOrder []int
}

func Constructor(root *TreeNode) BSTIterator {
	inOrder := make([]int, 0)
	bst(root, &inOrder)
	return BSTIterator{
		index:   0,
		inOrder: inOrder,
	}
}

func bst(root *TreeNode, inOrder *[]int) {
	if root == nil {
		return
	}
	bst(root.Left, inOrder)
	*inOrder = append(*inOrder, root.Val)
	bst(root.Right, inOrder)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	ret := this.inOrder[this.index]
	this.index++
	return ret
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.index < len(this.inOrder) {
		return true
	}
	return false
}
