package regularexpressionmatching

// https://leetcode-cn.com/problems/regular-expression-matching/

/*
Given:
T[] items = new T[width * height];

Converting 2D co-ordinates into 1D index:
y * width + x

Converting 1D index into 2D co-ordinates:
y = index / width;
x = index % width;

 */
func isMatch(s string, p string) bool {
	var (
		dp func(si, pi int) bool
	)

	x := len(s) + 1
	y := len(p) + 1
	memo := make([]int, x * y)

	dp = func(si, pi int) bool {
		if memo[si * y + pi] != 0 {
			return memo[si * y + pi] == 1
		}
		if pi == len(p) {
			return si == len(s)
		}

		firstMatch := (si < len(s)) && (p[pi] == s[si] || p[pi] == '.')

		if (pi <= len(p) - 2) && (p[pi + 1] == '*') {
			ans := dp(si, pi + 2) || (firstMatch && dp(si + 1, pi))
			if ans {
				memo[si * y + pi] = 1
			} else {
				memo[si * y + pi] = 2
			}
			return ans
		}
		ans := firstMatch && dp(si + 1, pi + 1)
		if ans {
			memo[si * y + pi] = 1
		} else {
			memo[si * y + pi] = 2
		}
		return ans
	}
	return dp(0, 0)
}