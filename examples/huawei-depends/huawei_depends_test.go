package huawei_depends

import (
	"fmt"
	"testing"
)

func TestAssignTaskOrder_Case1(t *testing.T) {
	ret := assignTaskOrder(0, [][]int{{4, 3}, {9, 8}, {3, 2}, {8, 7}, {2, 1}, {7, 6}})
	fmt.Println(ret)
}

func TestFinalMerge_Case1(t *testing.T) {
	finalMerge([][]int{{1, 2, 3}, {7, 8, 9}, {3, 4, 5, 6, 7}})
}
