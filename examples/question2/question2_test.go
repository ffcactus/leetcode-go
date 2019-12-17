package question2

import (
	"fmt"
	"testing"
)

func TestQuestion2_Case1(t *testing.T) {
	input := []int{1, 2, 2, 3, 1}
	output := findShortestDuration(input)
	fmt.Println(output)
}

func TestQuestion2_Case2(t *testing.T) {
	input := []int{1, 2, 2, 3, 1, 4, 2}
	output := findShortestDuration(input)
	fmt.Println(output)
}

func TestQuestion2_Case3(t *testing.T) {
	input := []int{1}
	output := findShortestDuration(input)
	fmt.Println(output)
}

func TestQuestion2_Case4(t *testing.T) {
	input := []int{1, 2}
	output := findShortestDuration(input)
	fmt.Println(output)
	input = []int{2, 1, 2}
	output = findShortestDuration(input)
	fmt.Println(output)
}

func TestQuestion2_Case5(t *testing.T) {
	input := []int{1, 2, 3, 4, 4, 3, 2, 1, 1, 4, 3, 2}
	output := findShortestDuration(input)
	fmt.Println(output)
}

func TestQuestion2_Case6(t *testing.T) {
	input := []int{2, 2, 2, 2}
	output := findShortestDuration(input)
	fmt.Println(output)
}

func TestQuestion2_Case7(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	output := findShortestDuration(input)
	fmt.Println(output)
}
