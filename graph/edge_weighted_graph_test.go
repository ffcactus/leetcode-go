package graph

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tinyEWG = `8
16
4 5 0.35
4 7 0.37
5 7 0.28
0 7 0.16
1 5 0.32
0 4 0.38
2 3 0.17
1 7 0.19
0 2 0.26
1 2 0.36
1 3 0.29
2 7 0.34
6 2 0.40
3 6 0.52
6 0 0.58
6 4 0.93
`

func TestNewEdgeWeightedGraphFromReader(t *testing.T) {
	g, err := NewEdgeWeightedGraphFromReader(strings.NewReader(tinyEWG))
	assert.NoError(t, err)
	impl, _ := g.(*adjacentListImplOfEWG)
	assert.ElementsMatch(t, []Edge{{v:6, w: 0, weight: 0.58}, {v:0, w: 2, weight: 0.26}, {v:0, w: 4, weight: 0.38}, {v:0, w: 7, weight: 0.16}}, impl.adj[0])
}
