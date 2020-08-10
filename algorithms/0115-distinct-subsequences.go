package algorithms

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for x := range dp {
		dp[x] = make([]int, len(t)+1)
	}
	for x := range dp {
		for y := range dp[x] {
			if y == 0 {
				dp[x][y] = 1
				continue
			}
			if x == 0 {
				dp[x][y] = 0
				continue
			}
			if s[x-1] == t[y-1] {
				dp[x][y] = dp[x-1][y-1] + dp[x-1][y]
			} else {
				dp[x][y] = dp[x-1][y]
			}
		}
	}
	return dp[len(s)][len(t)]
}
