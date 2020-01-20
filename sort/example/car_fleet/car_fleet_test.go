package car_fleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	assert.Equal(t, 3, carFleet(12, position, speed))
}

func TestCarFleet_Case2(t *testing.T) {
	position := []int{0, 4, 2}
	speed := []int{2, 1, 3}
	assert.Equal(t, 1, carFleet(10, position, speed))
}

func TestCarFleet_Case3(t *testing.T) {
	position := []int{6, 8}
	speed := []int{3, 2}
	assert.Equal(t, 2, carFleet(10, position, speed))
}
