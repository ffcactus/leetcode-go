package edgeweighted

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/graph"
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

func TestNewLazyPrimMST(t *testing.T) {
	g, err := graph.NewEdgeWeightedGraphFromReader(strings.NewReader(tinyEWG))
	assert.NoError(t, err)
	mst := NewLazyPrimMST(g)
	// [0-7 0.16 1-7 0.19 0-2 0.26 2-3 0.17 5-7 0.28 4-5 0.35 6-2 0.40]
	fmt.Println(mst.Edges())
}
