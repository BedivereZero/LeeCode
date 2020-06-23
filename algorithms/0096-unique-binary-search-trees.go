package algorithms

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	cache := make([]int, n + 1)
	cache[0], cache[1] = 1, 1
	for i := 2; i < n + 1; i++ {
		for j := 0; j < i; j++ {
			cache[i] += cache[j] * cache[i - 1 - j]
		}
	}
	return cache[n]
}
