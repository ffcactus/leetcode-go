package find_longest_word

import (
	"sort"
)

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) != len(d[j]) {
			return len(d[i]) > len(d[j])
		}
		return d[i] < d[j]
	})
	for _, v := range d {
		if match(s, v) {
			return v
		}
	}
	return ""
}

func match(s, d string) bool {
	j := 0
	for i := 0; i < len(s) && j < len(d); i++ {
		if d[j] == s[i] {
			j++
		}
	}
	return j == len(d)
}
