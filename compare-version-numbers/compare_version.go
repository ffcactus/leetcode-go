package compare_version_numbers

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/compare-version-numbers/

func compareVersion(version1 string, version2 string) int {
	vs1 := strings.Split(version1, ".")
	vs2 := strings.Split(version2, ".")
	if len(vs1) > len(vs2) {
		c := len(vs1) - len(vs2)
		for i:=0; i < c; i++ {
			vs2 = append(vs2, "0")
		}
	} else if len(vs2) > len(vs1) {
		c := len(vs2) - len(vs1)
		for i:=0; i < c; i++ {
			vs1 = append(vs1, "0")
		}
	}

	for i:=0; i <len(vs1); i++ {
		v1, _ := strconv.ParseInt(vs1[i], 10, 64)
		v2, _ := strconv.ParseInt(vs2[i], 10, 64)
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
	}
	return 0
}
