package pq

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple": 2,
		"pear": 4,
	}
	pq := make(PriorityQueue, len(items))
	i:=0
	for value, priority := range items {
		pq[i] = &Item{
			value: value,
			priority: priority,
			index: i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value: "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	// Then change the value of the item inserted, it's priority should change.
	pq.update(item, "watermelon", 5)

	// Now, begin to test.
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 5, item.priority)
	assert.Equal(t, "watermelon", item.value)
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 4, item.priority)
	assert.Equal(t, "pear", item.value)
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 3, item.priority)
	assert.Equal(t, "banana", item.value)
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 2, item.priority)
	assert.Equal(t, "apple", item.value)

	assert.Equal(t, 0, pq.Len())
}
