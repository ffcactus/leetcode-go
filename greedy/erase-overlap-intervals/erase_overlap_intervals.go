package eraseoverlapintervals

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	N := len(intervals)
	if N == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	end := intervals[0][1]
	for _, v := range intervals {
		if v[0] >= end {
			count++
			end = v[1]
		}
	}
	return N - count
}
