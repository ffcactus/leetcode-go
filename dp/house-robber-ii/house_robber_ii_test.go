package house_robber_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob_Case1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 8
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case2(t *testing.T) {
	input := []int{2, 3, 2}
	expected := 3
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case3(t *testing.T) {
	input := []int{1, 2, 3, 1}
	expected := 4
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case4(t *testing.T) {
	input := []int{1, 1, 3, 6, 7, 10, 7, 1, 8, 5, 9, 1, 4, 4, 3}
	expected := 4
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case5(t *testing.T) {
	input := []int{10}
	expected := 10
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case6(t *testing.T) {
	input := []int{10, 11}
	expected := 11
	actual := rob(input)
	assert.Equal(t, expected, actual)
}
