package sum_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNumbers(t *testing.T) {
	n5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	n9 := &TreeNode{
		Val:   9,
		Left:  n5,
		Right: n1,
	}
	n0 := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:  n9,
		Right: n0,
	}
	sum := sumNumbers(n4)
	assert.Equal(t, 1026, sum)
}
