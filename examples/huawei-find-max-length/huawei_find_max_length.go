package huawei_find_max_length

const (
	ST = 128
)

func findMaxLength(arr []string) int {
	sum := ""
	for _, v := range arr {
		sum += v
	}

}


/*
func findMaxLength(arr []string) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	for len(arr) != 1 {
		if hasDuplicateChar(arr[0], arr[1]) {
			if len(arr) > 2 {
				arr = append(arr[0:1], arr[2:]...)
			} else if len(arr) == 2 {
				return len(arr[0])
			}
		} else {
			arr[0] = arr[0] + arr[1]
			arr = append(arr[0:1], arr[2:]...)
		}
	}
	sta := st(arr[0])
	for _, v := range sta {
		if v > 1 {
			return 0
		}
	}

	return len(arr[0])
}
*/

func st(s string) []int {
	ret := make([]int, ST)
	for _, c := range s {
		i := int(c)
		ret[i]++
	}
	return ret
}

func hasDuplicateChar(a, b string) bool {
	sta := st(a)
	stb := st(b)
	for i := range sta {
		if sta[i] > 0 && stb[i] > 0 {
			return true
		}
	}
	return false
}
