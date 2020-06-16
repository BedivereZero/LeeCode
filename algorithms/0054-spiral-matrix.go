package algorithms

func spiralOrder(matrix [][]int) []int {
	return surround(matrix, 0)
}

func surround(matrix [][]int, level int) []int {
	lenx := len(matrix) - 2 * level
	if lenx < 1 {
		return []int {}
	}
	leny := len(matrix[0]) - 2 * level
	if leny < 1 {
		return []int {}
	}
	if lenx == 1 {
		return matrix[level][level:level + leny]
	}
	length := leny
	if lenx > 1 {
		length += lenx - 1
	}
	if leny > 1 {
		length += leny - 1
	}
	if lenx > 1 && leny > 1 {
		length += lenx - 2
	}
	result := make([]int, length)
	for i := 0; i < leny; i++ {
		result[i] = matrix[level][level + i]
	}
	if lenx > 1 {
		for i := 0; i + 1 < lenx; i++ {
			result[i + leny] = matrix[level + i + 1][level + leny - 1]
		}
	}
	if lenx > 1 {
		for i := 0; i + 1 < leny; i++ {
			result[i + lenx + leny - 1] = matrix[level + lenx - 1][level + leny - 2 - i]
		}
	}
	if leny > 1 {
		for i := 0; i + 2 < lenx; i++ {
			result[i + lenx + 2 * leny - 2] = matrix[level + lenx - 2 - i][level]
		}
	}
	return append(result, surround(matrix, level + 1)...)
}
