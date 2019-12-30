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
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its Priority.
	item := &Item{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)

	// Then change the Value of the item inserted, it's Priority should change.
	pq.update(item, "watermelon", 5)

	// Now, begin to test.
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 5, item.Priority)
	assert.Equal(t, "watermelon", item.Value)
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 4, item.Priority)
	assert.Equal(t, "pear", item.Value)
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 3, item.Priority)
	assert.Equal(t, "banana", item.Value)
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, 2, item.Priority)
	assert.Equal(t, "apple", item.Value)

	assert.Equal(t, 0, pq.Len())
}
