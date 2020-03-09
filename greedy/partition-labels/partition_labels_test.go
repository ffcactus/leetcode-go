package partitionlabels

import (
	"testing"
)

func TestPartitionLabels_Case1(t *testing.T) {
	actual := partitionLabels("ababcbacadefegdehijhklij")
	t.Errorf("%v", actual)
}
