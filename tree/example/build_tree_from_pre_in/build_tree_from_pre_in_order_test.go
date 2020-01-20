package build_tree_from_pre_in

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
            3
       9        20
   nil nil   15     7
*/
func TestBuildTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	n3 := buildTree(preOrder, inOrder)
	assert.Equal(t, 3, n3.Val)
	n9 := n3.Left
	n20 := n3.Right
	assert.NotNil(t, n9)
	assert.NotNil(t, n20)
	assert.Equal(t, 9, n9.Val)
	assert.Equal(t, 20, n20.Val)
	n15 := n20.Left
	n7 := n20.Right
	assert.NotNil(t, n15)
	assert.NotNil(t, n7)
	assert.Equal(t, 15, n15.Val)
	assert.Equal(t, 7, n7.Val)
}
