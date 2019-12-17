package swapnum

import (
	"fmt"
	"testing"
)

func TestQuestion3_Case1(t *testing.T) {
	input := []int{1, 6, 4, 10, 2, 5}
	output := leftNum(input)
	fmt.Println(output)
}

func TestQuestion3_Case2(t *testing.T) {
	input := []int{2, 4, 1, 3, 6}
	output := leftNum(input)
	fmt.Println(output)
}

func TestQuestion3_Case3(t *testing.T) {
	var input = make([]int, 100000)
	for i := 0; i < 100000; i++ {
		input[i] = 100000 - i
	}
	output := leftNum(input)
	fmt.Println(output)
}

func TestQuestion3_Case4(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	output := leftNum(input)
	fmt.Println(output)
}
