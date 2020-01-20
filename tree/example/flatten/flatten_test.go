package flatten

import "testing"

func TestFlatten(t *testing.T) {
	var (
		n3 = TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		n4 = TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		}
		n2 = TreeNode{
			Val:   2,
			Left:  &n3,
			Right: &n4,
		}
		n6 = TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		}
		n5 = TreeNode{
			Val:   5,
			Left:  nil,
			Right: &n6,
		}
		n1 = TreeNode{
			Val:   1,
			Left:  &n2,
			Right: &n5,
		}
	)
	flatten(&n1)
}
