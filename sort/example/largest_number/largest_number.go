package largest_number

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strNums := make([]string, 0)
	for _, v := range nums {
		strNums = append(strNums, strconv.Itoa(v))
	}

	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	if strNums[0] == "0" {
		return "0"
	}
	ret := ""
	for _, v := range strNums {
		ret += v
	}
	return ret
}
