package ewd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDijkstraShortestPath(t *testing.T) {
	g, err := NewEdgeWeightedDirectedGraphFromStream(strings.NewReader(tinyEWD))
	assert.NoError(t, err)
	solution := NewDijkstraShortestPath(g, 0)
	assert.True(t, solution.HasPathTo(6))
	assert.Equal(t, 0.26+0.34+0.39+0.52, solution.DistanceTo(6))
	for _, e := range solution.PathTo(6) {
		fmt.Println(e)
	}
}
