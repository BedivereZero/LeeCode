package algorithms

func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			sum += prices[i-1] - prices[0]
			prices[0] = prices[i]
		}
	}
	if prices[0] < prices[len(prices)-1] {
		sum += prices[len(prices)-1] - prices[0]
	}
	return sum
}
