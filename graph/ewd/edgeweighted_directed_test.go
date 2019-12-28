package ewd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tinyEWD = `8
15
4 5 0.35
5 4 0.35
4 7 0.37
5 7 0.28
7 5 0.28
5 1 0.32
0 4 0.38
0 2 0.26
7 3 0.39
1 3 0.29
2 7 0.34
6 2 0.40
3 6 0.52
6 0 0.58
6 4 0.93
`
)
func TestNewEdgeWeightedDirectedGraphFromStream(t *testing.T) {
	g, err := NewEdgeWeightedDirectedGraphFromStream(strings.NewReader(tinyEWD))
	assert.NoError(t, err)
	assert.Equal(t, 8, g.Vertices())
	assert.Equal(t, 15, g.Edges())
	assert.ElementsMatch(t, []DirectedEdge{{0, 2, 0.26}, {0, 4, 0.38}}, g.Adjacent(0))
}
