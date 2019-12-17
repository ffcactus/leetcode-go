package unique_paths_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths_Case1(t *testing.T) {
	actual := uniquePaths(3, 2)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestUniquePaths_Case2(t *testing.T) {
	actual := uniquePaths(51, 9)
	expected := 1916797311
	assert.Equal(t, expected, actual)
}

func TestUniquePaths_Case3(t *testing.T) {
	actual := uniquePaths(13, 23)
	expected := 548354040
	assert.Equal(t, expected, actual)
}
