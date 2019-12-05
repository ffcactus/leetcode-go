package merge

// merge Do the merge of a[low, min] and a[min+1, high].
// It assume a[low, min] and a[min+1, high] are ordered.
func merge(a []int, low, mid, high int) {
	var (
		i   = low
		j   = mid + 1
		aux = make([]int, len(a))
	)

	for k := low; k <= high; k++ {
		aux[k] = a[k]
	}

	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > high {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
