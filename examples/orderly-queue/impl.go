package orderly_queue

import "sort"

func orderlyQueue(S string, K int) string {
	if K == 1 {
		min := S
		for i := 0; i < len(S); i++ {
			newS := S[i:] + S[:i]
			if newS < min {
				min = newS
			}
		}
		return min
	}
	b := []byte(S)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
