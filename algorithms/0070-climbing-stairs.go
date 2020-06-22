package algorithms

func climbStairs(n int) int {
	a, b := 1, 1
	for i := 2; i < n + 1; i++ {
		a, b = a + b, a
	}
	return a
}
