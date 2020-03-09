package trigle

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Errorf("%v", generate(5))
}

func TestGetRow(t *testing.T) {
	t.Errorf("%v", getRow(3))
}
