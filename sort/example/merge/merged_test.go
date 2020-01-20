package merge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge_Empty(t *testing.T) {
	var empty [][]int
	ret := merge(empty)
	assert.Empty(t, ret)
}

func TestMerge(t *testing.T) {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ret := merge(intervals)
	fmt.Println(ret)
	assert.ElementsMatch(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, ret)
}
