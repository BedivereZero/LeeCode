package algorithms

func maxProfit(prices []int) int {
	forward := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		if profit := prices[i] - prices[i-1] + forward[i-1]; profit > 0 {
			forward[i] = profit
		} else {
			forward[i] = 0
		}
	}
	for i := 1; i < len(prices); i++ {
		if forward[i] < forward[i-1] {
			forward[i] = forward[i-1]
		}
	}
	backward := make([]int, len(prices))
	for i := len(prices) - 2; i >= 0; i-- {
		if profit := prices[i+1] + backward[i+1] - prices[i]; profit > 0 {
			backward[i] = profit
		} else {
			backward[i] = 0
		}
	}
	for i := len(prices) - 2; i >= 0; i-- {
		if backward[i] < backward[i+1] {
			backward[i] = backward[i+1]
		}
	}
	max := 0
	for i := 0; i < len(prices); i++ {
		if max < forward[i]+backward[i] {
			max = forward[i] + backward[i]
		}
	}
	return max
}
