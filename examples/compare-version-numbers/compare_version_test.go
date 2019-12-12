package compare_version_numbers

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCompareVersion_1(t *testing.T) {
	assert.Equal(t, compareVersion("0.1", "1.1"), -1)
}

func TestCompareVersion_2(t *testing.T) {
	assert.Equal(t, compareVersion("1.0.1", "1"), 1)
}

func TestCompareVersion_3(t *testing.T) {
	assert.Equal(t, compareVersion("7.5.2.4", "7.5.3"), -1)
}

func TestCompareVersion_4(t *testing.T) {
	assert.Equal(t, compareVersion("1.01", "1.001"), 0)
}

func TestCompareVersion_5(t *testing.T) {
	assert.Equal(t, compareVersion("1.0", "1.0.0"), 0)
}
