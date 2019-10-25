package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.Equal(t, true, contains1("BANC", "ABC"))
	assert.Equal(t, true, contains1("ABC", "ABC"))
	assert.Equal(t, true, contains1("ABC", "CBA"))
	assert.Equal(t, false, contains1("AC", "ACC"))
	assert.Equal(t, false, contains1("BCC", "CCC"))
}

func TestSolution1_1(t *testing.T) {
	assert.Equal(t, "BANC", solution1("ADOBECODEBANC", "ABC"))
}

func TestSolution1_2(t *testing.T) {
	assert.Equal(t, "", solution1("a", "aa"))
}

func TestSolution1_3(t *testing.T) {
	assert.Equal(t, "aa", solution1("aa", "aa"))
}