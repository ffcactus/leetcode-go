package list

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	n3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	reverseList(n1)
}
