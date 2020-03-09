package reconstructqueue

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] < people[j][0] {
			return false
		} else {
			return people[i][1] < people[j][1]
		}
	})
	res := make([][]int, 0)
	for _, v := range people {
		insert(&res, v)
	}
	return res
}

func insert(people *[][]int, single []int) {
	i := single[1]
	// a = append(a[:i], append([]T{x}, a[i:]...)...)
	*people = append((*people)[:i], append([][]int{single}, (*people)[i:]...)...)
}
