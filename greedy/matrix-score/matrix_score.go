package matrixscore

func matrixScore(A [][]int) int {
	Y := len(A)
	X := len(A[0])
	for {
		changed := false
		for y := 0; y < Y; y++ {
			if A[y][0] == 0 {
				changed = true
				for x := 0; x < X; x++ {
					if A[y][x] == 1 {
						A[y][x] = 0
					} else {
						A[y][x] = 1
					}
				}
			}
		}
		for x := 0; x < X; x++ {
			numsOfOne := 0
			for y := 0; y < Y; y++ {
				if A[y][x] == 1 {
					numsOfOne++
				}
			}
			if numsOfOne < (Y+1)/2 {
				changed = true
				for y := 0; y < Y; y++ {
					if A[y][x] == 1 {
						A[y][x] = 0
					} else {
						A[y][x] = 1
					}
				}
			}
		}
		if !changed {
			break
		}
	}
	return calculate(A)
}

func calculate(A [][]int) int {
	sum := 0
	Y := len(A)
	X := len(A[0])
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if A[y][x] == 1 {
				sum += (1 << uint64(X-x-1))
			}
		}
	}
	return sum
}
