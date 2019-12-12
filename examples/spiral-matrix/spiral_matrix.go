package spiralmatrix

// https://leetcode-cn.com/problems/spiral-matrix

const (
	right = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {
	maxy := len(matrix)
	if maxy == 0 {
		return []int{}
	}
	if maxy == 1 {
		return matrix[0]
	}
	maxx := len(matrix[0])
	upWall := 0
	rightWall := maxx
	downWall := maxy
	leftWall := -1
	ret := make([]int, 0)
	direction := right
	count := maxy * maxx
	x := 0
	y := 0
	for count > 0 {
		// fmt.Println(matrix[y][x])
		ret = append(ret, matrix[y][x])
		count--
		switch direction {
		case right:
			x++
			if x == rightWall {
				direction = down
				y++
				x--
				rightWall--
			}
		case down:
			y++
			if y == downWall {
				direction = left
				x--
				y--
				downWall--
			}
		case left:
			x--
			if x == leftWall {
				direction = up
				y--
				x++
				leftWall++
			}
		case up:
			y--
			if y == upWall {
				direction = right
				x++
				y++
				upWall++
			}
		}
	}
	return ret
}
