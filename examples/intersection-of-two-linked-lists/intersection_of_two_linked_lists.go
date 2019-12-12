package intersection_of_two_linked_lists

/**
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}

	lenA := lenOfListNode(headA)
	lenB := lenOfListNode(headB)
	if lenA > lenB {
		headA = move(lenA-lenB, headA)
	} else {
		headB = move(lenB-lenA, headB)
	}

	for true {
		if headA == headB {
			return headA
		}
		if headA == nil || headB == nil {
			return nil
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func lenOfListNode(node *ListNode) int {
	l := 0
	if node == nil {
		return l
	}
	for node.Next != nil {
		l++
		node = node.Next
	}
	return l
}

func move(count int, head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for ; count > 0; count-- {
		head = head.Next
	}
	return head
}
