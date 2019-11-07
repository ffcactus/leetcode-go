package swim_in_rising_water

import "sort"

type record struct {
	v int
	x int
	y int
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	surface := make([]record, 0)
	path := make([]record, 0)
	x := 0
	y := 0
	updateSurface(&surface, nil, []record{{
		v: grid[y][x],
		x: x,
		y: y,
	}})
	for true {
		current := surface[0]
		path = append(path, current)
		x = current.x
		y = current.y
		if (x == n-1) && (y == n-1) {
			break
		}
		newRecords := possibleNewRecord(current.x, current.y, n, path, grid)
		updateSurface(&surface, &current, newRecords)
	}

	max := 0
	for _, v := range path {
		if v.v > max {
			max = v.v
		}
	}
	return max
}

func possibleNewRecord(x, y, n int, path []record, grid [][]int) []record {
	var ret []record

	if y > 0 {
		up := record{
			x: x,
			y: y - 1,
			v: grid[y-1][x],
		}
		isOld := false
		for _, v := range path {
			if up.x == v.x && up.y == v.y {
				isOld = true
				break
			}
		}
		if !isOld {
			ret = append(ret, up)
		}
	}
	if y < n-1 {
		down := record{
			x: x,
			y: y + 1,
			v: grid[y+1][x],
		}
		isOld := false
		for _, v := range path {
			if down.x == v.x && down.y == v.y {
				isOld = true
				break
			}
		}
		if !isOld {
			ret = append(ret, down)
		}
	}
	if x > 0 {
		left := record{
			x: x - 1,
			y: y,
			v: grid[y][x-1],
		}
		isOld := false
		for _, v := range path {
			if left.x == v.x && left.y == v.y {
				isOld = true
				break
			}
		}
		if !isOld {
			ret = append(ret, left)
		}
	}
	if x < n-1 {
		right := record{
			x: x + 1,
			y: y,
			v: grid[y][x+1],
		}
		isOld := false
		for _, v := range path {
			if right.x == v.x && right.y == v.y {
				isOld = true
				break
			}
		}
		if !isOld {
			ret = append(ret, right)
		}
	}

	return ret
}

func updateSurface(surface *[]record, toRemove *record, newRecord []record) {
	// add new records to surface.
	*surface = append(*surface, newRecord...)
	// remove the toRemove from surface.
	if toRemove != nil {
		for i := range *surface {
			if ((*surface)[i].x == (*toRemove).x) && ((*surface)[i].y == (*toRemove).y) {
				*surface = append((*surface)[:i], (*surface)[i+1:]...)
				break
			}
		}
	}
	// sort the surface
	sort.Slice(*surface, func(i, j int) bool {
		if (*surface)[i].v < (*surface)[j].v {
			return true
		}
		return false
	})
}
