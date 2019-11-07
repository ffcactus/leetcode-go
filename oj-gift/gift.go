package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ParseIntBase    = 10
	ParseIntBitSize = 64
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		numOfColorString := scanner.Text()
		scanner.Scan()
		colorsString := scanner.Text()
		scanner.Scan()
		mString := scanner.Text()

		strconv.ParseInt(numOfColorString, ParseIntBase, ParseIntBitSize) //

		colorsStrings := strings.Split(colorsString, " ")
		colors := make([]int64, 0)
		for _, v := range colorsStrings {
			vv, _ := strconv.ParseInt(v, ParseIntBase, ParseIntBitSize)
			colors = append(colors, vv)
		}
		m, _ := strconv.ParseInt(mString, ParseIntBase, ParseIntBitSize)
		count := solution1(colors, m)
		fmt.Println(count)
	}
}

func solution1(colors []int64, m int64) int64 {
	fmt.Println(len(colors))
	fmt.Println(m)
	sum := int64(0)
	for _, v := range colors {
		sum += v
	}
	max := sum / m
	for max > 0 {
		validPearls := int64(0)
		l := len(colors)
		for i := 0; i < l; i++ {
			if colors[i] > max {
				colors[i] = max
			}
			validPearls += colors[i]
		}
		if validPearls >= max*m {
			break
		}
		max = validPearls / m
	}
	return max
}

func solution2(colors []int, m int) int {
	// sort.Ints 按从小大大排列
	sort.Ints(colors)
	l := len(colors)
	for i := 0; i < l/2; i++ {
		colors[i], colors[l-i-1] = colors[l-i-1], colors[i]
	}
	// 记录最后输出。
	count := 0
	for len(colors) >= m {
		// 从头开始取m个珠子。
		success := true
		i := 0
		for i = 0; i < m; i++ {
			colors[i]--
			if colors[i] < 0 {
				success = false
				break
			}
		}
		if !success {
			break
		}
		count++
		// 判断是否需要重新排列。当取完M个珠子后，排在后面的珠子比前面的大时需要重新排列。
		if i < len(colors) {
			if colors[i-1] < colors[i] {
				// 重新排列下
				// sort.Ints(colors)
				// l := len(colors)
				// for i:=0; i < l / 2; i++ {
				// 	colors[i], colors[l - i - 1] = colors[l - i - 1], colors[i]
				// }
				colors = merge(colors[:i], colors[i:])
			}
		}
		// 清除出已经为0的珠子。
		newColors := make([]int, 0)
		for i := 0; i < len(colors); i++ {
			if colors[i] > 0 {
				newColors = append(newColors, colors[i])
			}
		}
		colors = newColors
	}
	return count
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}
