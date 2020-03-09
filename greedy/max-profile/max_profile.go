package maxprofile

func maxProfit(prices []int, fee int) int {
	N := len(prices)
	cash := 0
	hold := -prices[0]
	for i := 1; i < N; i++ {
		cash = max2(cash, hold+prices[i]-fee)
		hold = max2(hold, cash-prices[i])
	}
	return cash
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
