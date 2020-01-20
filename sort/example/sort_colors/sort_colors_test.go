package sort_colors

import "testing"

func TestSortColors(t *testing.T) {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

func TestSortColors_1(t *testing.T) {
	sortColors([]int{2, 0, 1})
}
