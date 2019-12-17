package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/graph"
)

var testStream = `13
13
0 5
4 3
0 1
9 12
6 4
5 4
0 2
11 12
9 10
0 6
7 8
9 11
5 3
`

func TestNewDepthFirstSearch(t *testing.T) {
	g, err := graph.NewUndirectedGraphFromReader(strings.NewReader(testStream))
	assert.NoError(t, err)
	// check 1
	impl := NewDepthFirstSearch(g, 0)
	marks := make([]int, 0)
	for i := 0; i < g.Vertices(); i++ {
		if impl.Marked(i) {
			marks = append(marks, i)
		}
	}
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6}, marks)

	// check 2
	impl = NewDepthFirstSearch(g, 9)
	marks = make([]int, 0)
	for i := 0; i < g.Vertices(); i++ {
		if impl.Marked(i) {
			marks = append(marks, i)
		}
	}
	assert.ElementsMatch(t, []int{9, 10, 11, 12}, marks)
}
