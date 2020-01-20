package wiggle_sort

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	input := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(input)
	fmt.Println(input)
}
