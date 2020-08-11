package algorithms

func maxProfit(prices []int) int {
	max := 0
	profit := make([]int, len(prices))
	for s := 0; s < len(prices); s++ {
		// calculate first half
		max0 := 0
		for i := 1; i < s; i++ {
			if prices[i] < prices[i-1]-profit[i-1] {
				profit[i] = 0
				continue
			}
			profit[i] = prices[i] - prices[i-1] + profit[i-1]
			if max0 < profit[i] {
				max0 = profit[i]
			}
		}
		// calculate second half
		max1 := 0
		for i := s + 1; i < len(prices); i++ {
			if prices[i] < prices[i-1]-profit[i-1] {
				profit[i] = 0
				continue
			}
			profit[i] = prices[i] - prices[i-1] + profit[i-1]
			if max1 < profit[i] {
				max1 = profit[i]
			}
		}
		// clean cache
		for i := 0; i < len(prices); i++ {
			profit[i] = 0
		}
		if max < max0+max1 {
			max = max0 + max1
		}
	}
	return max
}
