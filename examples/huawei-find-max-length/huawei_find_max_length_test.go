package huawei_find_max_length

import (
	"fmt"
	"testing"
)

func TestFindMaxLength_Case1(t *testing.T) {
	input := []string{"un","iq","ue"}
	actual := findMaxLength(input)
	fmt.Println(actual)
}

func TestFindMaxLength_Case2(t *testing.T) {
	input := []string{"aabcdefghijklmnopqrstuvwxyz"}
	actual := findMaxLength(input)
	fmt.Println(actual)
}
