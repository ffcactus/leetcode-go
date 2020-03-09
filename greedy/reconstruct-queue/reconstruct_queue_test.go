package reconstructqueue

import (
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	input := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	actual := reconstructQueue(input)
	t.Errorf("%v", actual)
}
