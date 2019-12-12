package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	actual := spiralOrder(matrix)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, actual = %v", expected, actual)
	}
}
