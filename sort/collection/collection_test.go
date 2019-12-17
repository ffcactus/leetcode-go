package collection

import (
	"sort"
	"testing"
)

func TestStringCollection_Case1(t *testing.T) {
	sc := StringCollection{raw: []string{"ccc", "bbb", "aaa"}}
	sort.Sort(sc)
	if !sort.IsSorted(sc) {
		t.Errorf("Failed.")
	}
}
