package algorithms
func generateMatrix(n int) [][]int {
	// generate matrix, n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// fill it
	for level, begin := 0, 1; level < (n + 1) / 2; level++ {
		generateEdge(matrix, level, begin)
		begin += (n - 1 - level * 2) * 4
	}

	return matrix
}

func generateEdge(matrix [][]int, level int, begin int) {
	length := len(matrix) - 2 * level
	if length == 1 {
		matrix[level][level] = begin
		return
	}
	for i := 0; i + 1 < length; i++ {
		matrix[level][level + i] = begin + i
	}
	for i := 0; i + 1 < length; i++ {
		matrix[level + i][level + length - 1] = begin + length - 1 + i
	}
	for i := 0; i + 1 < length; i++ {
		matrix[level + length - 1][level + length - 1 - i] = begin + (length - 1) * 2 + i
	}
	for i := 0; i + 1 < length; i++ {
		matrix[level + length - 1 - i][level] = begin + (length - 1) * 3 + i
	}
}
