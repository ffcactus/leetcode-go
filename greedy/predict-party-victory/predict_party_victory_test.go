package predictpartyvictory

import (
	"testing"
)

func TestPredictPartyVictory_Case1(t *testing.T) {
	actual := predictPartyVictory("RDD")
	expected := "Dire"
	if actual != expected {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}

func TestPredictPartyVictory_Case2(t *testing.T) {
	actual := predictPartyVictory("RD")
	expected := "Radiant"
	if actual != expected {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}
