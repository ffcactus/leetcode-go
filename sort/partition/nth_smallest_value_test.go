package partition

import "testing"

func TestNthSmallest_Case0(t *testing.T) {
	a := []int{5, 4, 3, 2, 1, 0}
	if actual := nthSmallest(a, 3); actual != 3 {
		t.Errorf("Expected 3 but is %d", actual)
	}
}

func TestNthSmallest_Case1(t *testing.T) {
	a := []int{5, 4, 3, 2, 1, 0}
	if actual := nthSmallest(a, 0); actual != 0 {
		t.Errorf("Expected 3 but is %d", actual)
	}
}