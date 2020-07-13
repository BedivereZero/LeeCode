package algorithms

func maximalRectangle(matrix [][]byte) int {

	maximum := 0

	for x := range matrix {
		// Monotonic stack
		index := make([]int, len(matrix[x]))
		cache := make([]int, len(matrix[x]))
		depth := 0

		for y := range matrix[x] {
			// transform to histogram
			if x > 0 && matrix[x][y] == '1' {
				matrix[x][y] = matrix[x-1][y] + 1
			}

			for depth > 0 && matrix[x][index[depth-1]] > matrix[x][y] {
				depth--
				if s := int(matrix[x][index[depth]]-'0') * (y - index[depth] - 1); maximum < s+cache[depth] {
					maximum = s + cache[depth]
				}
			}

			// width
			w := y + 1
			if depth > 0 {
				w -= index[depth-1] + 1
			}
			// square
			s := w * int(matrix[x][y]-'0')
			if depth > 0 && matrix[x][index[depth-1]] == matrix[x][y] {
				depth--
				s += cache[depth]
			}
			index[depth] = y
			cache[depth] = s
			depth++

			if maximum < s {
				maximum = s
			}
		}

		for depth > 0 {
			depth--
			if s := int(matrix[x][index[depth]]-'0') * (len(matrix[x]) - index[depth] - 1); maximum < s+cache[depth] {
				maximum = s + cache[depth]
			}
		}
	}
	return maximum
}
