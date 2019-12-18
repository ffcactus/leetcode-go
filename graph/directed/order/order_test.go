package order

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode-go/graph"
	"strings"
	"testing"
)

var stream = `13
15
0 1
0 5
0 6
2 0
2 3
3 5
5 4
6 4
6 9
7 6
8 7
9 10
9 11
9 12
11 12
`

func TestNewDFSOrder(t *testing.T) {
	g, err := graph.NewDirectedGraphFromReader(strings.NewReader(stream))
	assert.NoError(t, err)
	impl := NewDFSOrder(g)
	fmt.Println(impl.Pre())
	fmt.Println(impl.Post())
	fmt.Println(impl.ReversePost())
}
