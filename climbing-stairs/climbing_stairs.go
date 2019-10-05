package climbingstairs

func climbStairs(n int) int {
	var memo = make([]int, n + 1)
	return _climbStairs(0, n, &memo)
}

func _climbStairs(i, n int, memo *[]int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1;
	}
	if (*memo)[i] > 0 {
		return (*memo)[i]
	}
	(*memo)[i] = _climbStairs(i + 1, n, memo) + _climbStairs(i + 2, n, memo)
	return (*memo)[i]
}