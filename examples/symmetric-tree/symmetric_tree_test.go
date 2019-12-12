package symmetric_tree

import (
	"testing"
)

func TestSymmetricTree_SingleNode(t *testing.T) {
	a1 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	actual := isSymmetric(&a1)
	expected := true
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestSymmetricTree_2Node(t *testing.T) {
	a2 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	a1 := TreeNode{
		Val:   4,
		Left:  &a2,
		Right: nil,
	}
	actual := isSymmetric(&a1)
	expected := false
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestSymmetricTree_NormalSymmetricTree(t *testing.T) {
	a4 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	a5 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	a6 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	a7 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	a2 := TreeNode{
		Val:   2,
		Left:  &a4,
		Right: &a5,
	}
	a3 := TreeNode{
		Val:   2,
		Left:  &a6,
		Right: &a7,
	}

	a1 := TreeNode{
		Val:   1,
		Left:  &a2,
		Right: &a3,
	}

	actual := isSymmetric(&a1)
	expected := true
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestSymmetricTree_NotSymmetricTree(t *testing.T) {
	d3 := TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	d4 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	d7 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	d8 := TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	c1 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	c2 := TreeNode{
		Val:   5,
		Left:  &d3,
		Right: &d4,
	}
	c3 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	c4 := TreeNode{
		Val:   4,
		Left:  &d7,
		Right: &d8,
	}
	b1 := TreeNode{
		Val:   3,
		Left:  &c1,
		Right: &c2,
	}
	b2 := TreeNode{
		Val:   3,
		Left:  &c3,
		Right: &c4,
	}
	a1 := TreeNode{
		Val:   2,
		Left:  &b1,
		Right: &b2,
	}

	actual := isSymmetric(&a1)
	expected := false
	if actual != expected {
		t.Errorf("Failed.")
	}
}
