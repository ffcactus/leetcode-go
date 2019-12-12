package remove_invalid_parentheses

// https://leetcode-cn.com/problems/remove-invalid-parentheses/

func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)

	// 记录多出来的左右括号。
	bracketLeft := 0
	bracketRight := 0
	for _, v := range s {
		if v == '(' {
			bracketLeft++
		}
		if v == ')' {
			if bracketLeft > 0 {
				bracketLeft--
			} else {
				bracketRight++
			}
		}
	}

	dfs(s, 0, bracketLeft, bracketRight, &ans)
	return ans
}

// dfs 从start开始，对s进行深度优先遍历。
func dfs(s string, start, bracketLeft, bracketRight int, ans *[]string) {
	if bracketLeft == 0 && bracketRight == 0 {
		if valid(s) {
			*ans = append(*ans, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i-1 >= start && s[i] == s[i-1] {
			continue
		}
		if bracketLeft > 0 && s[i] == '(' {
			b := []byte(s)
			bs := append(b[:i], b[i+1:]...)
			dfs(string(bs), i, bracketLeft-1, bracketRight, ans)
		}
		if bracketRight > 0 && s[i] == ')' {
			b := []byte(s)
			bs := append(b[:i], b[i+1:]...)
			dfs(string(bs), i, bracketLeft, bracketRight-1, ans)
		}
	}
}

// valid 判断s是不是合规的表达式。
func valid(s string) bool {
	cnt := 0
	for _, v := range s {
		if v == '(' {
			cnt++
		}
		if v == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
