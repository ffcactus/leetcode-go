package move_zeroes

import (
	"fmt"
	"testing"
)

func TestMoveZeroes_0(t *testing.T) {
	nums := []int{0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestMoveZeroes_1(t *testing.T) {
	nums := []int{0, 1, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestMoveZeroes_2(t *testing.T) {
	nums := []int{0,1,0,3,4}
	moveZeroes(nums)
	fmt.Println(nums)
}
