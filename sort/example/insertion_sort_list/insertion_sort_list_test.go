package insertion_sort_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortList_Case0(t *testing.T) {
	assert.Nil(t, insertionSortList(nil))
}

func TestInsertionSortList_Case1(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	head = insertionSortList(head)
	assert.Equal(t, 1, head.Val)
	head = head.Next
	assert.Equal(t, 2, head.Val)
	head = head.Next
	assert.Equal(t, 3, head.Val)
	head = head.Next
	assert.Equal(t, 4, head.Val)
	head = head.Next
	assert.Nil(t, head)
}
