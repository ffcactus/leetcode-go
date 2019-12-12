package regularexpressionmatching

import (
	"testing"
)

func TestIsMatch_Case1(t *testing.T) {
	actual := isMatch("aa", "a")
	expected := false
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case2(t *testing.T) {
	actual := isMatch("", "a")
	expected := false
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case3(t *testing.T) {
	actual := isMatch("a", "")
	expected := false
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case4(t *testing.T) {
	actual := isMatch("aa", "a*")
	expected := true
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case5(t *testing.T) {
	actual := isMatch("ab", ".*")
	expected := true
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case6(t *testing.T) {
	actual := isMatch("aab", "c*a*b")
	expected := true
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case7(t *testing.T) {
	actual := isMatch("mississippi", "mis*is*p*.")
	expected := false
	if actual != expected {
		t.Errorf("Failed.")
	}
}

func TestIsMatch_Case8(t *testing.T) {
	actual := isMatch("aaa", "a*a")
	expected := true
	if actual != expected {
		t.Errorf("Failed.")
	}
}
