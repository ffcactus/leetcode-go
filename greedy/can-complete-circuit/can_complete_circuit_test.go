package circuit

import (
	"testing"
)

func TestCanCompleteCircuit_Case1(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	result := canCompleteCircuit(gas, cost)
	if result != 3 {
		t.Errorf("Expect = %d, actual is %d", 3, result)
	}
}

func TestCanCompleteCircuit_Case2(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	result := canCompleteCircuit(gas, cost)
	if result != -1 {
		t.Errorf("Expect = %d, actual is %d", -1, result)
	}
}
