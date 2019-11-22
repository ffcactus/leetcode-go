package swapnum

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	l := lenOfList(head)
	left := toIndex(k-1, head)
	right := toIndex(l-k, head)
	left.Val, right.Val = right.Val, left.Val
	return head
}

func toIndex(k int, head *ListNode) *ListNode {
	var ret *ListNode = head
	for k > 0 {
		ret = ret.Next
		k--
	}
	return ret
}

func lenOfList(head *ListNode) int {
	if head == nil {
		return 0
	}
	ret := 0
	for head.Next != nil {
		ret++
		head = head.Next
	}
	ret++
	return ret
}
