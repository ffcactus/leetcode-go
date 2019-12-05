package merge

import (
	"sort"
	"testing"
)

func TestMerge_Case1(t *testing.T) {
	a := []int{100}
	merge(a, 0, 0, 0)
}

func TestMerge_Case4(t *testing.T) {
	a := []int{2,1}
	merge(a, 0, 0, 1)
	if !sort.IntsAreSorted(a) {
		t.Errorf("array is not sorted.")
	}
}


func TestMerge_Case2(t *testing.T) {
	a := []int{1, 3, 5, 7, 2, 4, 6, 8}
	merge(a, 0, 3, 7)
	if !sort.IntsAreSorted(a) {
		t.Errorf("array is not sorted.")
	}
}

func TestMerge_Case3(t *testing.T) {
	a := []int{1, 3, 5, 2, 4, 6, 8, 10}
	merge(a, 0, 2, 7)
	if !sort.IntsAreSorted(a) {
		t.Errorf("array is not sorted.")
	}
}
