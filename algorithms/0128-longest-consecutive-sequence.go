package algorithms

func longestConsecutive(nums []int) int {
	// Endpoints -> Phase
	ep := make(map[int]int)
	longest := 0
	for _, n := range nums {
		if _, ok := ep[n]; ok || ep[n-1] > 0 || ep[n+1] < 0 {
			continue
		}
		l, r := n, n
		if p, ok := ep[n-1]; ok {
			l = n - 1 + ep[n-1]
			if p < 0 {
				delete(ep, n-1)
			}
		}
		if p, ok := ep[n+1]; ok {
			r = n + 1 + ep[n+1]
			if p > 0 {
				delete(ep, n+1)
			}
		}
		ep[l], ep[r] = r-l, l-r
		if longest < r-l+1 {
			longest = r - l + 1
		}
	}
	return longest
}
