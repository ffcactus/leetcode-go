package contains_nearby_almost_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	assert.False(t, containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
