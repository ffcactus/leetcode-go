package pairs

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	n4 := ListNode{
		Val:  4,
		Next: nil,
	}
	n3 := ListNode{
		Val:  3,
		Next: &n4,
	}
	n2 := ListNode{
		Val:  2,
		Next: &n3,
	}
	n1 := ListNode{
		Val:  1,
		Next: &n2,
	}
	swapPairs(&n1)
	t.Errorf("")
}
