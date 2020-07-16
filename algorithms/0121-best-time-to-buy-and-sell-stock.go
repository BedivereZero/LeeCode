package algorithms

func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if max < prices[i]-prices[i-1] {
			max = prices[i] - prices[i-1]
		}
		if prices[i-1] < prices[i] {
			prices[i] = prices[i-1]
		}
	}
	return max
}
