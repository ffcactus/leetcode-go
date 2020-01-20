package car_fleet

import "sort"

type car struct {
	position int
	t        float64
}

func carFleet(target int, position []int, speed []int) int {
	N := len(position)

	if N == 0 || N == 1 {
		return N
	}
	cars := make([]car, N)
	for i := range position {
		cars[i].position = position[i]
		cars[i].t = float64(target-position[i]) / float64(speed[i])
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position < cars[j].position
	})
	ans := 0
	i := N
	for i > 1 {
		i--
		if cars[i].t < cars[i-1].t {
			ans++
		} else {
			cars[i-1] = cars[i]
		}
	}
	ans++
	return ans
}
