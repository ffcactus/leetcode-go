package longest_consecutive_sequence

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCase1(t *testing.T) {
	input := []int{}
	expected := 0
	actual := longestConsecutive(input)
	assert.Equal(t, actual, expected, "Case 1 failed.")
}

func TestCase2(t *testing.T) {
	input := []int{5}
	expected := 1
	actual := longestConsecutive(input)
	assert.Equal(t, actual, expected, "Case 2 failed.")
}

func TestCase3(t *testing.T) {
	input := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	actual := longestConsecutive(input)
	assert.Equal(t, actual, expected, "Case 3 failed.")
}

func TestCase4(t *testing.T) {
	input := []int{1, 2, 0, 1}
	expected := 3
	actual := longestConsecutive(input)
	assert.Equal(t, actual, expected, "Case 4 failed.")
}
