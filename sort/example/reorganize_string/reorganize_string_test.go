package reorganize_string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorganizeString_Case1(t *testing.T) {
	assert.Equal(t, "", reorganizeString(""))
}

func TestReorganizeString_Case2(t *testing.T) {
	fmt.Println(reorganizeString("aab"))
}

func TestReorganizeString_Case3(t *testing.T) {
	fmt.Println(reorganizeString("aaab"))
}

func TestReorganizeString_Case4(t *testing.T) {
	fmt.Println(reorganizeString("aaaaabcde"))
}
