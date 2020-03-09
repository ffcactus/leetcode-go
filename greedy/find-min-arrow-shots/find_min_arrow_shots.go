package findminarrowshots

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	N := len(points)
	if N == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	end := points[0][1]
	for _, v := range points {
		if v[0] > end {
			count++
			end = v[1]
		}
	}
	return count
}
