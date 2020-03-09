package length

import (
	"testing"
)

func TestWiggleMaxLength_Case1(t *testing.T) {
	input := []int{1, 7, 4, 9, 2, 5}
	expected := 6
	actual := wiggleMaxLength(input)
	if actual != expected {
		t.Errorf("Expected = %d, Actual = %d", expected, actual)
	}
}

func TestWiggleMaxLength_Case2(t *testing.T) {
	input := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	expected := 7
	actual := wiggleMaxLength(input)
	if actual != expected {
		t.Errorf("Expected = %d, Actual = %d", expected, actual)
	}
}

func TestWiggleMaxLength_Case3(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 2
	actual := wiggleMaxLength(input)
	if actual != expected {
		t.Errorf("Expected = %d, Actual = %d", expected, actual)
	}
}

func TestWiggleMaxLength_Case4(t *testing.T) {
	input := []int{3, 3, 3, 2, 5}
	expected := 3
	actual := wiggleMaxLength(input)
	if actual != expected {
		t.Errorf("Expected = %d, Actual = %d", expected, actual)
	}
}
