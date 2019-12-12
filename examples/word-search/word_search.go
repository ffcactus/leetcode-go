package word_search

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	maxY := len(board)
	if maxY == 0 {
		return false
	}
	maxX := len(board[0])
	if maxX == 0 {
		return false
	}
	wordBytes := []byte(word)
	if maxY*maxX < len(wordBytes) {
		return false
	}

	// boardStatistic := make(map[int]int)
	// wordStatistic := make(map[int]int)
	//
	// for _, v := range wordBytes {
	// 	wordStatistic[int(v)]++
	// }
	// for x := 0; x < maxX; x++ {
	// 	for y := 0; y < maxY; y++ {
	// 		boardStatistic[int(board[y][x])]++
	// 	}
	// }
	// for k, v := range wordStatistic {
	// 	if boardStatistic[k] < v {
	// 		fmt.Println("statistic take effect")
	// 		return false
	// 	}
	// }

	marker := createMarker(maxX, maxY)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {

			if foundFrom(x, y, maxX, maxY, board, 0, wordBytes, marker) {
				return true
			}
		}
	}
	return false
}

func createMarker(maxX, maxY int) [][]bool {
	ret := make([][]bool, maxY)
	for i := range ret {
		ret[i] = make([]bool, maxX)
	}
	return ret
}

// foundFrom 假定之前的寻找已经成功，并判断从[y][x]board开始，判断接下来的上下左右4个位置上是否有[wi]word这个字符。
func foundFrom(x, y, maxX, maxY int, border [][]byte, wi int, word []byte, marker [][]bool) bool {
	// 如果Word已经找完则返回。
	if wi >= len(word) {
		return true
	}
	// 在没找完的情况下。

	// 如果已经超出 board 范围则返回失败。
	if x >= maxX || y >= maxY || x < 0 || y < 0 {
		return false
	}

	if marker[y][x] {
		return false
	}
	if border[y][x] != word[wi] {
		return false
	}
	marker[y][x] = true
	wi++
	// 如果接下来的内容能在上方找到则找到了。
	if foundFrom(x, y-1, maxX, maxY, border, wi, word, marker) {
		return true
	}
	if foundFrom(x, y+1, maxX, maxY, border, wi, word, marker) {
		return true
	}
	if foundFrom(x-1, y, maxX, maxY, border, wi, word, marker) {
		return true
	}
	if foundFrom(x+1, y, maxX, maxY, border, wi, word, marker) {
		return true
	}
	// 如果上下左右都找不到，那么这次的起始位置也就没有意义了，该路径不行。
	marker[y][x] = false
	return false
}
