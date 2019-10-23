package invalid_transations

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/invalid-transactions/


type transaction struct {
	name string
	time int
	value int
	city string
	invalid bool
}

func (t transaction) String() string {
	return fmt.Sprintf("%s,%d,%d,%s", t.name, t.time, t.value, t.city)
}

func invalidTransactions(transactions []string) []string {
	// 先将原始数据转换成便于使用的内部数据。
	trans := make([]transaction, 0)
	for _, v := range transactions {
		trans = append(trans, parseTransaction(v))
	}

	// 在按名称排序的基础上再按交易时间排序。
	sort.Slice(trans, func(i, j int) bool {
		if trans[i].name < trans[j].name {
			return true
		}
		if trans[i].name > trans[j].name {
			return false
		}
		return trans[i].time < trans[j].time
	})

	for i := range trans {
		// if trans[i].invalid {
		// 	continue
		// }
		for j := i + 1; j < len(trans); j++ {
			// if trans[j].invalid {
			// 	continue
			// }
			// 如果名称不一样了，则没有必要继续。
			if trans[i].name != trans[j].name {
				break
			}
			if trans[j].time - trans[i].time <= 60 && trans[i].city != trans[j].city {
				trans[i].invalid = true
				trans[j].invalid = true
			}
			// 如果时间已经超过了 60 分钟则没有必要比较下去
			if trans[j].time - trans[i].time > 60 {
				break
			}
		}
	}

	// 将 invalid输出。
	ret := make([]string, 0)
	for _, v := range trans {
		if v.invalid {
			ret = append(ret, v.String())
		}
	}
	return ret
}


func parseTransaction(s string) transaction {
	ss := strings.Split(s, ",")
	t, _ := strconv.ParseInt(ss[1], 10, 64)
	v, _ := strconv.ParseInt(ss[2], 10, 64)
	return transaction{
		name:  ss[0],
		time:  int(t),
		value: int(v),
		city:  ss[3],
		invalid: v > 1000,
	}
}
