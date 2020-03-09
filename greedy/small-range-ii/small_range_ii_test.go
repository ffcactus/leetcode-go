package smallrangeii

import (
	"testing"
)

func TestSmallRangeII_Case1(t *testing.T) {
	actual := smallestRangeII([]int{1}, 0)
	expected := 0
	if actual != expected {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}

func TestSmallRangeII_Case2(t *testing.T) {
	actual := smallestRangeII([]int{0, 10}, 2)
	expected := 6
	if actual != expected {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}

func TestSmallRangeII_Case3(t *testing.T) {
	actual := smallestRangeII([]int{1, 3, 6}, 3)
	expected := 3
	if actual != expected {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}

func TestSmallRangeII_Case4(t *testing.T) {
	actual := smallestRangeII([]int{10, 7, 1}, 3)
	expected := 3
	if actual != expected {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}
