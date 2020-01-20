package level_order

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/tree/example/flatten"
)

func TestLevelOrder(t *testing.T) {
	n15 := &flatten.TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	n7 := &flatten.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n20 := &flatten.TreeNode{
		Val:   20,
		Left:  n15,
		Right: n7,
	}
	n9 := &flatten.TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	n3 := &flatten.TreeNode{
		Val:   3,
		Left:  n9,
		Right: n20,
	}
	ret := levelOrder(n3)
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, ret)
}

func TestLevelOrder_Empty(t *testing.T) {
	ret := levelOrder(nil)
	assert.Empty(t, ret)
}
