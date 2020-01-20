package is_valid_bst

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/tree/example/flatten"
)

func TestIsValidBST(t *testing.T) {
	var (
		n3 = &flatten.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		n6 = &flatten.TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		}
		n4 = &flatten.TreeNode{
			Val:   4,
			Left:  n3,
			Right: n6,
		}
		n1 = &flatten.TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}
		n5 = &flatten.TreeNode{
			Val:   5,
			Left:  n1,
			Right: n4,
		}
	)

	assert.False(t, isValidBST(n5))
}
