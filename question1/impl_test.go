package swapnum

import "testing"

func TestCase1(t *testing.T) {
	node6 := &ListNode {
		Val: 6,
	}
	node5 := &ListNode {
		Val: 5,
		Next: node6,
	}
	node4 := &ListNode {
		Val: 4,
		Next: node5,
	}
	node3 := &ListNode {
		Val: 3,
		Next: node4,
	}
	node2 := &ListNode {
		Val: 2,
		Next: node3,
	}
	node1 := &ListNode {
		Val: 1,
		Next: node2,
	}
	swap(node1, 2)
}
