package main

import (
	"testing"
)

func TestSolution1_1(t *testing.T) {
	colors := []int{4, 4, 3, 3, 3, 1}
	actual := solution1(colors, 3)
	expected := 6
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}

func TestSolution1_2(t *testing.T) {
	colors := []int{1, 2, 3, 4, 5, 6}
	actual := solution1(colors, 5)
	expected := 3
	if expected != actual {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}


func TestMerge(t *testing.T) {
	a := []int{5, 3, 1}
	b := []int{6, 4, 2}
	actual := merge(a, b)
	t.Error(actual)
}