package unique_paths_two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles_Case1(t *testing.T) {
	obstaclesGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	actual := uniquePathsWithObstacles(obstaclesGrid)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestUniquePathsWithObstacles_Case2(t *testing.T) {
	obstaclesGrid := [][]int{{1}}
	actual := uniquePathsWithObstacles(obstaclesGrid)
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestUniquePathsWithObstacles_Case3(t *testing.T) {
	obstaclesGrid := [][]int{{1, 0}}
	actual := uniquePathsWithObstacles(obstaclesGrid)
	expected := 0
	assert.Equal(t, expected, actual)
}