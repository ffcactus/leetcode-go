package grammer

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if N == 2 {
		if K == 1 {
			return 0
		}
		return 1
	}
	pre := kthGrammar(N-1, (K+1)/2)
	if K%2 == 0 {
		if pre == 0 {
			return 1
		}
		return 0
	}
	if pre == 0 {
		return 0
	}
	return 1
}
