package directed

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode-go/graph"
)

var stream = `4
4
0 1
1 2
2 3
3 1
`

func TestNewCycleDetection(t *testing.T) {
	g, err := graph.NewDirectedGraphFromReader(strings.NewReader(stream))
	assert.NoError(t, err)
	detector := NewCycleDetection(g)
	assert.True(t, detector.HasCycle())
	assert.Equal(t, []int{3, 1, 2, 3}, detector.Cycle())
}
