package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Mid return the middle node of the list.
func Mid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
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

func sortList(head *ListNode) *ListNode {
	return MergeSort(head)
}
