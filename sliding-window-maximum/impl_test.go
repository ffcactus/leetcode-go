package sliding_window_maximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow(input, 3), "should match")
}
