package house_robber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob_Case1(t *testing.T) {
	input := []int{1,2,3,1}
	expected := 4
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case2(t *testing.T) {
	input := []int{2,7,9,3,1}
	expected := 12
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case3(t *testing.T) {
	input := []int{2, 1, 1, 2}
	expected := 4
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case4(t *testing.T) {
	input := []int{9}
	expected := 9
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case5(t *testing.T) {
	input := []int{9, 10}
	expected := 10
	actual := rob(input)
	assert.Equal(t, expected, actual)
}

func TestRob_Case6(t *testing.T) {
	input := []int{1, 9, 2}
	expected := 9
	actual := rob(input)
	assert.Equal(t, expected, actual)
}