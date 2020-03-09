package splitintofibonacci

import (
	"strconv"
	"strings"
)

const (
	max = (2<<31 - 1)
)

func splitIntoFibonacci(S string) []int {
	N := len(S)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i > N || i+j > N {
				continue
			}
			first := S[:i]
			second := S[i : i+j]
			if first[0] == '0' && len(first) > 1 {
				continue
			}
			if second[0] == '0' && len(second) > 1 {
				continue
			}
			a, _ := strconv.Atoi(first)
			b, _ := strconv.Atoi(second)
			if a > max || b > max {
				continue
			}
			if ret := isFibonacci(S[len(first)+len(second):], a, b); ret != nil {
				return ret
			}
		}
	}
	return nil
}

func isFibonacci(S string, first, second int) []int {
	if len(S) == 0 {
		return nil
	}
	ret := make([]int, 0)
	ret = append(ret, first, second)

	for {
		next := first + second
		if next > max {
			return nil
		}
		nextS := strconv.Itoa(next)
		if strings.HasPrefix(S, nextS) {
			S = S[len(nextS):]
			ret = append(ret, next)
			first = second
			second = next
		} else if len(S) == 0 {
			return ret
		} else {
			return nil
		}
	}
}
