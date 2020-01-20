package dfs

import "testing"

/*
				1
		2    			3
4			5	6			7
*/
func TestDFSPost(t *testing.T) {
	n7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:  n6,
		Right: n7,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:  n4,
		Right: n5,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:  n2,
		Right: n3,
	}
	dfsPost(n1)
}
