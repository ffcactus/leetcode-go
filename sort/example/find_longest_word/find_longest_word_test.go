package find_longest_word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestWord_Case1(t *testing.T) {
	assert.Equal(t, "a", findLongestWord("abpcplea", []string{"a", "b", "c"}))
}

func TestFindLongestWord_Case2(t *testing.T) {
	assert.Equal(t, "apple", findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}

func TestFindLongestWord_Case3(t *testing.T) {
	assert.Equal(t, "ewaf", findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
}
