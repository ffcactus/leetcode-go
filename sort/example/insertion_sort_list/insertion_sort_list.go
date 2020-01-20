package insertion_sort_list

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return strconv.Itoa(l.Val)
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}

	// let current be the list that is need to be sorted.
	current := head.Next
	// let head be the list that is already sorted.
	head.Next = nil

	for current != nil {
		p := fakeHead
		for p.Next != nil && p.Next.Val < current.Val {
			p = p.Next
		}
		// now p is the previous place to insert.
		tmp := current.Next
		current.Next = p.Next
		p.Next = current
		current = tmp
	}
	return fakeHead.Next
}
