package merge

import "sort"

// merge merges the intervals and output the intervals that only includes all the covered area.
func merge(intervals [][]int) [][]int {
	// first sort the intervals by the start point.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var merged [][]int
	if len(intervals) == 0 {
		return merged
	}
	// put the first one into the merged output.
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastMerged := &merged[len(merged)-1]
		// if the start of the current interval bigger than the end of previous interval, we can put the current one into the merged output.
		if intervals[i][0] > (*lastMerged)[1] {
			merged = append(merged, intervals[i])
			continue
		}
		// else they must have overlapped area.
		// In this case , we extend the end of the previous interval to the value of max(previous.end, current.end)
		if (*lastMerged)[1] < intervals[i][1] {
			(*lastMerged)[1] = intervals[i][1]
		}
	}
	return merged
}
