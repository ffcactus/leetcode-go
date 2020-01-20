package h_index

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHIndex(t *testing.T) {
	assert.Equal(t, 3, hIndex([]int{3, 0, 6, 1, 5}))
}
