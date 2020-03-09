package maxprofile

import (
	"testing"
)

func TestMaxProfile(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	actual := maxProfit(prices, fee)
	expected := 8

	if actual != expected {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}
