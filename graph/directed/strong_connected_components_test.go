package directed

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/graph"
)

var tinyDG = `13
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

func TestNewKosarajuSSCChecker(t *testing.T) {
	g, err := graph.NewDirectedGraphFromReader(strings.NewReader(tinyDG))
	assert.NoError(t, err)
	checker := NewKosarajuSSCChecker(g)
	assert.Equal(t, 5, checker.Count())
}
