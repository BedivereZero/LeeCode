package algorithms

func trap(height []int) int {
	// pool
	pool := make([]int, len(height))

	// calculate level
	for i := range pool {
		// pool height and left position
		h, l := 0, i

		// find left edge of pool
		for j := i - 1; h < height[i] && j >= 0; j-- {
			if h < height[j] {
				h = height[j]
				l = j
			}
			if height[i] < h {
				h = height[i]
			}
		}

		// update level in the pool
		for j := l + 1; j < i; j++ {
			pool[j] = h
		}

		pool[i] = height[i]
	}
	cap := 0
	for i := range height {
		cap += pool[i] - height[i]
	}
	return cap
}
