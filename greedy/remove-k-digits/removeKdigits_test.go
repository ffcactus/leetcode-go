package removekdigits

import (
	"testing"
)

func TestRemoveKDigits_Case0(t *testing.T) {
	expected := "1219"
	actual := removeKdigits("1432219", 3)
	if expected != actual {
		t.Errorf("Expected = %s, actual = %s", expected, actual)
	}
}

func TestRemoveKDigits_Case1(t *testing.T) {
	expected := "200"
	actual := removeKdigits("10200", 1)
	if expected != actual {
		t.Errorf("Expected = %s, actual = %s", expected, actual)
	}
}

func TestRemoveKDigits_Case2(t *testing.T) {
	expected := "0"
	actual := removeKdigits("10", 2)
	if expected != actual {
		t.Errorf("Expected = %s, actual = %s", expected, actual)
	}
}
