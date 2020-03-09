package trigle

func generate(numRows int) [][]int {
	var (
		i   int
		ret [][]int
	)
	if numRows == 0 {
		return ret
	}
	ret = append(ret, []int{1})
	i = 1
	for i < numRows {
		next := make([]int, i+1)
		for j := range next {
			if j == 0 {
				next[j] = 1
			} else if j == len(next)-1 {
				next[j] = 1
			} else {
				next[j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
		ret = append(ret, next)
		i++
	}
	return ret
}

func getRow(rowIndex int) []int {
	return helper(rowIndex + 1)
}

func helper(rowIndex int) []int {
	if rowIndex == 1 {
		return []int{1}
	}
	pre := helper(rowIndex - 1)
	cur := make([]int, rowIndex)
	for j := range cur {
		if j == 0 {
			cur[j] = 1
		} else if j == len(cur)-1 {
			cur[j] = 1
		} else {
			cur[j] = pre[j-1] + pre[j]
		}
	}
	return cur
}
