package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	resetHead, _ := helper(head)
	return resetHead
}

func helper(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return head, head
	}
	if head.Next == nil {
		head.Next = head
		return head, head
	}
	resetHead, resetTail := helper(head.Next)
	resetTail.Next = head
	head.Next = nil
	return resetHead, head
}
