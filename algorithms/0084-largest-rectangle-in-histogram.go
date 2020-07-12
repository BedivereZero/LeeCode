package algorithms

func largestRectangleArea(heights []int) int {
	index := make([]int, len(heights))
	cache := make([]int, len(heights))
	depth := 0

	largest := 0
	for i, h := range heights {
		for depth > 0 && heights[index[depth-1]] > h {
			depth--
			if s := heights[index[depth]] * (i - index[depth] - 1); largest < s+cache[depth] {
				largest = s + cache[depth]
			}
		}
		width := i + 1
		if depth > 0 {
			width -= index[depth-1] + 1
		}
		s := width * h
		if depth > 0 && heights[index[depth-1]] == h {
			s += cache[depth-1]
			depth--
		}
		index[depth] = i
		cache[depth] = s
		depth++
		if largest < s {
			largest = s
		}
	}
	for depth > 0 {
		depth--
		if s := heights[index[depth]] * (len(heights) - index[depth] - 1); largest < s+cache[depth] {
			largest = s + cache[depth]
		}
	}
	return largest
}
