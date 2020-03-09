package numrescureboats

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	light := 0
	heavy := len(people) - 1
	count := 0
	for light <= heavy {
		if people[heavy]+people[light] <= limit {
			count++
			light++
			heavy--
		} else {
			count++
			heavy--
		}
	}
	return count
}
