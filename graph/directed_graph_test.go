package graph

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var directedGraphInputStream = `13
22
4 2
2 3
3 2
6 0
0 1
2 0
11 12
12 9
9 10
9 11
8 9
10 12
11 4
4 3
3 5
7 8
8 7
5 4
0 5
6 4
6 9
7 6
`

// TestNewUndirectedGraphFromReader check if the adjacent lists in the implementation match expectation.
func TestNewDirectedGraphFromReader(t *testing.T) {
	dg, err := NewDirectedGraphFromReader(strings.NewReader(directedGraphInputStream))
	if err != nil {
		t.Errorf("Should not have error but got %v", err)
	}
	assert.NoError(t, err)
	impl, _ := dg.(*directedGraphImpl)
	assert.ElementsMatch(t, []int{5, 1}, impl.adj[0])
	assert.ElementsMatch(t, []int{}, impl.adj[1])
	assert.ElementsMatch(t, []int{0, 3}, impl.adj[2])
	assert.ElementsMatch(t, []int{5, 2}, impl.adj[3])
	assert.ElementsMatch(t, []int{3, 2}, impl.adj[4])
	assert.ElementsMatch(t, []int{4}, impl.adj[5])
	assert.ElementsMatch(t, []int{9, 4, 0}, impl.adj[6])
	assert.ElementsMatch(t, []int{6, 8}, impl.adj[7])
	assert.ElementsMatch(t, []int{7, 9}, impl.adj[8])
	assert.ElementsMatch(t, []int{11, 10}, impl.adj[9])
	assert.ElementsMatch(t, []int{12}, impl.adj[10])
	assert.ElementsMatch(t, []int{4, 12}, impl.adj[11])
	assert.ElementsMatch(t, []int{9}, impl.adj[12])
}
