func maxArea(height []int) int {
	a, l, r := 0, 0, len(height) - 1
	for l < r {
		var h int
		if height[l] < height[r] {
			h = height[l]
		} else {
			h = height[r]
		}
		ca := h * (r - l)
		if a < ca {
			a = ca
		}
		for l < len(height) && height[l] <= h {
			l++
		}
		for r >= 0 && height[r] <= h {
			r --
		}
	}
	return a
}
