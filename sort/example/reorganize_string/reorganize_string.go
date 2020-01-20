package reorganize_string

import "sort"

type element struct {
	alpha byte
	count int
}

func reorganizeString(S string) string {
	if S == "" {
		return ""
	}

	b := []byte(S)
	st := make([]element, 26)
	for i := range st {
		st[i].alpha = byte('a' + i)
	}
	for _, v := range b {
		st[v-'a'].count++
	}
	sort.Slice(st, func(i, j int) bool {
		return st[i].count > st[j].count
	})
	i := st[0].count
	if len(S)%2 == 1 && i > len(S)/2+1 {
		return ""
	}
	if len(S)%2 == 0 && i > len(S)/2 {
		return ""
	}
	i = 0
	for len(st) > 0 {
		t := st[0].alpha
		k := st[0].count
		for k > 0 {
			k--
			if i >= len(S) {
				i = 1
			}
			b[i] = t
			i += 2
		}
		st = st[1:]
	}
	S = string(b)
	return S
}
