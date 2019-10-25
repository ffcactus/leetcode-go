package minimum_window_substring

func minWindow(s string, t string) string {
	return solution1(s, t)
}

func solution1(s string, t string) string {
	candidate := make([]string, 0)
	if t == "" {
		return ""
	}
	left, right := 0, len(t)
	// sizeS, sizeT := len(s), len(t)
	for right <= len(s) {
		if contains1(s[left:right], t) {
			// reduce the window by increase the left.
			for right-left >= len(t) {
				left++
				if !contains1(s[left:right], t) {
					// fmt.Printf("Add candidate %s\n", s[left - 1:right])
					candidate = append(candidate, s[left-1:right])
					break
				}
			}
		}
		right++
	}
	if len(candidate) == 0 {
		return ""
	}
	ret := s
	for _, v := range candidate {
		if len(v) < len(ret) {
			ret = v
		}
	}
	return ret
}

func contains(s, t string) bool {
	// fmt.Printf("Checkt s = %s, t = %s\n", s, t)
	copyS := []byte(s)
	copyT := []byte(t)
	for i := range copyT {
		found := false
		for j := range copyS {
			if copyS[j] == copyT[i] {
				copyS[j] = 0
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func contains1(s, t string) bool {
	arrayS := make([]int, 127)
	arrayT := make([]int, 127)
	for _, v := range s {
		arrayS[v]++
	}
	for _, v := range t {
		arrayT[v]++
	}
	for i:=0; i < 127; i++ {
		if arrayS[i] < arrayT[i] {
			return false
		}
	}
	return true
}
