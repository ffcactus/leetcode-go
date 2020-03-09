package leastinterval

import (
	"testing"
)

func TestLeastInterval_Case1(t *testing.T) {
	b := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	expected := 8
	actual := leastInterval(b, n)
	if expected != actual {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}

func TestLeastInterval_Case2(t *testing.T) {
	b := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'}
	n := 2
	expected := 9
	actual := leastInterval(b, n)
	if expected != actual {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}

func TestLeastInterval_Case3(t *testing.T) {
	b := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2
	expected := 16
	actual := leastInterval(b, n)
	if expected != actual {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}
