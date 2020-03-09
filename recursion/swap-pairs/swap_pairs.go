package pairs

// ListNode defines the element.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	t0 := head
	t1 := head.Next.Next
	head = head.Next
	head.Next = t0
	head.Next.Next = swapPairs(t1)
	return head
}
