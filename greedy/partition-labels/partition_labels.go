package partitionlabels

func partitionLabels(S string) []int {
	tails := make([]int, 26)
	for i := 0; i < len(S); i++ {
		tails[S[i]-'a'] = i
	}
	from := 0
	lastPosition := 0
	ans := make([]int, 0)
	for i := 0; i < len(S); i++ {
		lastPosition = max(lastPosition, tails[S[i]-'a'])
		if i == lastPosition {
			ans = append(ans, i-from+1)
			from = i + 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
