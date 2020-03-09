package circuit

func canCompleteCircuit(gas []int, cost []int) int {
	N := len(gas)
	for i := 0; i < N; i++ {
		good := true
		from := i
		currentGas := 0
		for j := 0; j < N; j++ {
			currentGas += gas[from]
			// check if we can reach to the next station.
			currentGas -= cost[from]
			if currentGas < 0 {
				good = false
				break
			} else {
				from++
				if from == N {
					from = 0
				}
			}

		}
		if good {
			return i
		}
	}
	return -1
}
