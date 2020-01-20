package h_index

import "sort"

func hIndex(citations []int) int {
	l := len(citations)
	if l == 0 {
		return 0
	}
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	ret := citations[0]
	for ret > 0 {
		found := true
		if ret <= l {
			for i := 0; i < ret; i++ {
				if citations[i] < ret {
					found = false
					break
				}
			}
		} else {
			found = false
		}

		if found {
			return ret
		}
		ret--
	}
	return ret
}
