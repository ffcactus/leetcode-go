package inorder_tree_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/tree/example/flatten"
)

func TestInorderTreeTraversal(t *testing.T) {
	var (
		n3 = &flatten.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		n2 = &flatten.TreeNode{
			Val:   2,
			Left:  n3,
			Right: nil,
		}
		n1 = &flatten.TreeNode{
			Val:   1,
			Left:  nil,
			Right: n2,
		}
	)

	ret := inorderTraversal(n1)
	assert.Equal(t, []int{1, 3, 2}, ret)
}
