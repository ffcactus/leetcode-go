package removekdigits

func removeKdigits(num string, k int) string {
	for i := 0; i < k; i++ {
		num = removeSingle(num)
	}
	return num
}

func removeSingle(num string) string {
	N := len(num)
	if N == 1 || N == 0 {
		return "0"
	}
	for i := 1; i < N; i++ {
		if num[i] < num[i-1] {
			result := num[:i-1] + num[i:]
			return removeHeadZero(result)
		}
	}
	return removeHeadZero(num[:N-1])
}

func removeHeadZero(num string) string {
	N := len(num)
	if N < 2 {
		return num
	}
	if num[0] != '0' {
		return num
	}
	count := 1
	for i := 1; i < N-1; i++ {
		if num[i] == '0' {
			count++
		} else {
			break
		}
	}
	return num[count:]
}
