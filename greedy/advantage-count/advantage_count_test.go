package advantagecount

import (
	"testing"
)

func TestAdvantageCount_Case1(t *testing.T) {
	A := []int{2, 7, 11, 15}
	B := []int{1, 10, 4, 11}
	actual := advantageCount(A, B)
	t.Errorf("%v", actual)
}

func TestAdvantageCount_Case2(t *testing.T) {
	A := []int{12, 24, 8, 32}
	B := []int{13, 25, 32, 11}
	actual := advantageCount(A, B)
	t.Errorf("%v", actual)
}
