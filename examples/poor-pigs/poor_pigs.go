package poor_pigs

import (
	"math"
)

// https://leetcode-cn.com/problems/poor-pigs/

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	turns := minutesToTest/minutesToDie + 1
	ret := math.Log(float64(buckets)) / math.Log(float64(turns))
	return int(math.Ceil(ret))
}
