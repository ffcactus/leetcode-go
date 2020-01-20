package common

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverse the order of node list.
func ReverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode
		curr *ListNode
	)
	curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// Mid return the middle node of the list.
func Mid(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}
	mid := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}
	return mid
}

// HasCycle check if there is a cycle in the list.
// Suppose we have 2 people running along the list, and one of the people running faster then the other.
// If there is a cycle the faster one will eventually meet the slower one in the middle of the cycle.
// If there is no cycle, the faster one will reach the end of the list.
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

// DetectCycle return the cycle head if there is.
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	start := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for start != slow {
				slow = slow.Next
				start = start.Next
			}
			return slow
		}
	}
	return nil
}

// Merge merges sorted list l1 and l2 into a sorted list.
func Merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = Merge(l1.Next, l2)
		return l1
	}
	l2.Next = Merge(l1, l2.Next)
	return l2
}

// MergeSort use merge sort algorithm to sort the list.
func MergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := Mid(head)
	right := mid.Next
	mid.Next = nil
	left := head
	return Merge(MergeSort(left), MergeSort(right))
}
