package monotoneincreasingdigits

import (
	"strconv"
)

func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	s := strconv.Itoa(N)
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		if b[i] > b[i+1] {
			b[i]--
			for j := i + 1; j < len(b); j++ {
				b[j] = '9'
			}
			i = -1
		}
	}
	ret, _ := strconv.Atoi(string(b))
	return ret
}

func helper(n int) int {
	s := strconv.Itoa(n)
	b := []byte(s)
	if len(b) > 1 {
		for i := 1; i < len(b); i++ {
			if b[i] < b[i-1] {
				for j := i; j < len(b); j++ {
					b[j] = '0'
				}
				break
			}
		}
	}
	ret, _ := strconv.Atoi(string(b))
	return ret
}
