package connectedcomponents

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

func TestNewDFSImpl(t *testing.T) {
	g, err := graph.NewUndirectedGraphFromReader(strings.NewReader(testStream))
	assert.NoError(t, err)
	cc := NewDFSImpl(g)
	assert.Equal(t, 3, cc.Count())
}
