package partition

import "testing"

func TestPartition_Case1(t *testing.T) {
	a := []int{4, 5, 3, 1, 2}
	N := len(a)
	k := partition(a, 0, 4)
	for i := k - 1; i >= 0; i-- {
		if a[i] > a[k] {
			t.Errorf("Left ones failed.")
			return
		}
	}

	for i := k + 1; i < N; i++ {
		if a[i] < a[k] {
			t.Errorf("Right ones failed.")
			return
		}
	}
}
