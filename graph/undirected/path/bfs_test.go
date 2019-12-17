package path

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/graph"
)

var tinyInput = `6
8
0 5
2 4
2 3
1 2
0 1
3 4
3 5
0 2
`

func TestNewBFSPath(t *testing.T) {
	g, err := graph.NewUndirectedGraphFromReader(strings.NewReader(tinyInput))
	assert.NoError(t, err)
	impl := NewBFSPath(g, 0)
	assert.True(t, impl.HasPathTo(0))
	assert.True(t, impl.HasPathTo(1))
	assert.True(t, impl.HasPathTo(2))
	assert.True(t, impl.HasPathTo(3))
	assert.True(t, impl.HasPathTo(4))
	assert.True(t, impl.HasPathTo(5))

	assert.True(t, sameOrder([]int{0}, impl.PathTo(0)))
	assert.True(t, sameOrder([]int{0, 1}, impl.PathTo(1)))
	assert.True(t, sameOrder([]int{0, 2}, impl.PathTo(2)))
	assert.True(t, sameOrder([]int{0, 2, 3}, impl.PathTo(3)))
	assert.True(t, sameOrder([]int{0, 2, 4}, impl.PathTo(4)))
	assert.True(t, sameOrder([]int{0, 5}, impl.PathTo(5)))
}

func sameOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
