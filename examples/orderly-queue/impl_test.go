package orderly_queue

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCase1(t *testing.T) {
	s := "aaacaaab"
	k := 1
	expected := "aaabaaac"
	actual := orderlyQueue(s, k)
	assert.Equal(t, actual, expected, "Case 1 failed.")
}

func TestCase2(t *testing.T) {
	s := "cba"
	k := 1
	expected := "acb"
	actual := orderlyQueue(s, k)
	assert.Equal(t, actual, expected, "Case 2 failed.")
}

func TestCase3(t *testing.T) {
	s := "baaca"
	k := 2
	expected := "aaabc"
	actual := orderlyQueue(s, k)
	assert.Equal(t, actual, expected, "Case 3 failed.")
}