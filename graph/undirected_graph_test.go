package graph

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var undirectedGraphInputStream = `13
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

// TestNewUndirectedGraphFromReader check if the adjacent lists in the implementation match expectation.
func TestNewUndirectedGraphFromReader(t *testing.T) {
	g, err := NewUndirectedGraphFromReader(strings.NewReader(undirectedGraphInputStream))
	assert.NoError(t, err)
	impl := g.(*undirectedGraphImpl)
	assert.ElementsMatch(t, []int{6, 2, 1, 5}, impl.adj[0])
	assert.ElementsMatch(t, []int{0}, impl.adj[1])
	assert.ElementsMatch(t, []int{0}, impl.adj[2])
	assert.ElementsMatch(t, []int{5, 4}, impl.adj[3])
	assert.ElementsMatch(t, []int{5, 6, 3}, impl.adj[4])
	assert.ElementsMatch(t, []int{3, 4, 0}, impl.adj[5])
	assert.ElementsMatch(t, []int{0, 4}, impl.adj[6])
	assert.ElementsMatch(t, []int{8}, impl.adj[7])
	assert.ElementsMatch(t, []int{7}, impl.adj[8])
	assert.ElementsMatch(t, []int{11, 10, 12}, impl.adj[9])
	assert.ElementsMatch(t, []int{9}, impl.adj[10])
	assert.ElementsMatch(t, []int{9, 12}, impl.adj[11])
	assert.ElementsMatch(t, []int{11, 9}, impl.adj[12])
}
