package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	// please define the input here.
	// For example: r := bufio.NewReader(os.Stdin) input, err := r.ReadString('\n')

	// please finish the function body here.

	// please define the output here. For example: fmt.Println(input)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ns := scanner.Text()
		n, _ := strconv.ParseInt(ns, 10, 64)
		fmt.Println(solution2(int(n)))
	}
}

// solution1 是纯数学解法。
// 结果等于 （3的n次方 + 1）/ 2
func solution1(n int) *big.Int {
	big3, bign := big.NewInt(3), big.NewInt(int64(n))
	big3.Exp(big3, bign, nil)
	big3.Add(big3, big.NewInt(1))
	big3.Div(big3, big.NewInt(2))
	return big3
}

// solution2 是迭代解法。
// 假设在第n个位置放置的是白色，为了满足白色为偶数，那么n之前的白色需要是奇数个。
// 假设在第n个位置放置的是红色，为了满足白色为偶数，那么n之前的白色需要是偶数个。
// 假设在第n个位置放置的是绿色，为了满足白色为偶数，那么n之前的白色需要是偶数个。
// 所以n个位置，白色为偶数的方案为上面3个之和。
// 同理也可以推算出n个位置，白色为奇数的方案数。
// 当得知n=2个位置时，白色为偶数的方案和白色为奇数的方案，那么就可以推理出n>2个位置时白色为偶数（或奇数）的方案。
func solution2(n int) *big.Int {

	// 1个方块时，白色为奇数的方案数，就等于当前放的是白色。
	preWhiteOdd := big.NewInt(1)
	// 1个方块时，白色为偶的方案数。白色为偶数则只可能没有白色，剩下放红和放绿。
	preWhiteEven := big.NewInt(2)
	i := 2
	for i <= n {
		// 算上当前位置，白色为奇数的方案: (当最后一个放白)preWhiteEvent + (当最后一个放红)preWhiteOdd + (当最后一个放绿)preWhiteOdd
		nowOdd := big.NewInt(0)
		nowOdd.Add(nowOdd, preWhiteEven)
		nowOdd.Add(nowOdd, preWhiteOdd)
		nowOdd.Add(nowOdd, preWhiteOdd)

		// 算上当前位置，白色为偶数的方案: (当最后一个放白)preWhiteOdd + (当最后一个放红)preWhiteEven + (当最后一个放绿)preWhiteEvent
		nowEven := big.NewInt(0)
		nowEven.Add(nowEven, preWhiteOdd)
		nowEven.Add(nowEven, preWhiteEven)
		nowEven.Add(nowEven, preWhiteEven)

		i++
		preWhiteOdd = nowOdd
		preWhiteEven = nowEven
	}
	return preWhiteEven
}
