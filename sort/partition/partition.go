package partition

// partition rearrange the array from index low to hight and return a index.
// that ones to the left of the index are no greater and the ones to the right of the index
// are no smaller
func partition(a []int, low, high int) int {
	var (
		i = low
		j = high + 1
		v = a[low]
	)

	for {
		for {
			i++
			if a[i] < v {
				if i == high {
					break
				}
			} else {
				break
			}
		}

		for {
			j--
			if v < a[j] {
				if j == low {
					break
				}
			} else {
				break
			}
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}

	a[low], a[j] = a[j], a[low]
	return j
}
