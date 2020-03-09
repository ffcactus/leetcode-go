package predictpartyvictory

func predictPartyVictory(senate string) string {
	queue := make([]int, 0)
	people := []int{0, 0}
	bans := []int{0, 0}

	for _, p := range senate {
		x := 0
		if p == 'R' {
			x = 1
		}
		people[x]++
		queue = append(queue, x)
	}
	for people[0] > 0 && people[1] > 0 {
		x := queue[0]
		queue = queue[1:]
		if bans[x] > 0 {
			bans[x]--
			people[x]--
		} else {
			bans[x^1]++
			queue = append(queue, x)
		}
	}
	if people[1] > 0 {
		return "Radiant"
	}
	return "Dire"
}
