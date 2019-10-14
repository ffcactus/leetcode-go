package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// main 函数解决该问题。
// 简单的来讲这是一个字符串排序的问题（将数字看成字符串）。
// 情形1，字符串长度相同。此时字符串大的应该在前面。
// 情形2，字符串长度不同，例如987和98, 那么98应该在前面，但是129和12，则129应该在前面。
// 为了解决情形2，可以直接将两字符串合并，在前在后各合并一次，然后比较合并后字符串的大小。比如129和12可以合并成12912和12129，再比较。
// 由于OJ上的Go版本较老，没有sort.Slice的函数，不能定制化比较函数，这里使用快速排序法，然后定制化其中的比较过程。
func main() {
	// please define the input here.
	// For example: r := bufio.NewReader(os.Stdin) input, err := r.ReadString('\n')

	// please finish the function body here.

	// please define the output here. For example: fmt.Println(input)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cases := scanner.Text()
	casesNum, _ := strconv.ParseInt(cases, 10, 64)
	for i:=0; i <int(casesNum); i++ {
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		numListString := scanner.Text()
		// fmt.Printf("num list string = %s\n", numListString)
		numList := strings.Split(numListString, " ")
		sortStrings(numList)
		for _, v := range numList {
			fmt.Printf("%s", v)
		}
		fmt.Printf("\n")
	}
}

func sortStrings(a []string) []string {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		// 定制化的比较过程。
		if a[i] + a[right] > a[right] + a[i] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sortStrings(a[:left])
	sortStrings(a[left+1:])
	return a
}