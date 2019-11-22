package unique_binary_search_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTrees_Case1(t *testing.T) {
	expect := 1
	actual := numTrees(1)
	assert.Equal(t, expect, actual)
}

func TestNumTrees_Case2(t *testing.T) {
	expect := 2
	actual := numTrees(2)
	assert.Equal(t, expect, actual)
}

func TestNumTrees_Case3(t *testing.T) {
	expect := 5
	actual := numTrees(3)
	assert.Equal(t, expect, actual)
}