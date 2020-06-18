package algorithms

func uniquePaths(m int, n int) int {
	// Make sure m <= n
	if n < m {
		m, n = n, m
	}
	return combination(m + n - 2, m - 1)
}

func combination(n int, k int) int {
	base := 1
	for i := 0; i < k; i++ {
		base = base * (n - i) / (i + 1)
	}
	return base
}
