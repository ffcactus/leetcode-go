package eraseoverlapintervals

import (
	"testing"
)

func TestEraseOverlapIntervals_Case0(t *testing.T) {
	input := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	expected := 1
	actual := eraseOverlapIntervals(input)
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestEraseOverlapIntervals_Case2(t *testing.T) {
	input := [][]int{{1, 2}, {1, 3}, {1, 2}}
	expected := 2
	actual := eraseOverlapIntervals(input)
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestEraseOverlapIntervals_Case3(t *testing.T) {
	input := [][]int{{1, 2}, {2, 3}}
	expected := 0
	actual := eraseOverlapIntervals(input)
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestEraseOverlapIntervals_Case4(t *testing.T) {
	input := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 3}, {3, 5}, {1, 5}}
	expected := 3
	actual := eraseOverlapIntervals(input)
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestEraseOverlapIntervals_Case5(t *testing.T) {
	input := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	expected := 2
	actual := eraseOverlapIntervals(input)
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestEraseOverlapIntervals_Case6(t *testing.T) {
	input := [][]int{{0, 2}, {1, 3}, {1, 3}, {2, 4}, {3, 5}, {3, 5}, {4, 6}}
	expected := 4
	actual := eraseOverlapIntervals(input)
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}
