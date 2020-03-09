package matrixscore

import (
	"testing"
)

func TestMatrixScore_Case1(t *testing.T) {
	input := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	matrixScore(input)
}
