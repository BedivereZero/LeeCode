package algorithms

func trap(height []int) int {
	// capcity
	cap := 0

	// calculate level
	for i := range height {
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
			cap += h - height[j]
			height[j] = h
		}
	}
	return cap
}
