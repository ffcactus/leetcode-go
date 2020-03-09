package ispossible

func isPossible(nums []int) bool {
	frequence := make(map[int]int)
	hyper := make(map[int]int)
	for _, v := range nums {
		frequence[v]++
	}

	for _, v := range nums {
		if frequence[v] == 0 {
			continue
		}

		if hyper[v] > 0 { // if num can be put into the chain.
			hyper[v]--
			hyper[v+1]++
			frequence[v]--
		} else if frequence[v+1] > 0 && frequence[v+2] > 0 {
			frequence[v]--
			frequence[v+1]--
			frequence[v+2]--
			hyper[v+3]++
		} else {
			return false
		}
	}
	return true
}
