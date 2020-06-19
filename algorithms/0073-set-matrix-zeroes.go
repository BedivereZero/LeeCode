package algorithms

func setZeroes(matrix [][]int) {
	recordX := make([]bool, len(matrix))
	recordY := make([]bool, len(matrix[0]))
	for x := range matrix {
		for y := range matrix[0] {
			if matrix[x][y] == 0 {
				recordX[x] = true
				recordY[y] = true
			}
		}
	}
	for index, value := range recordX {
		if value {
			for y := range matrix[0] {
				matrix[index][y] = 0
			}
		}
	}
	for index, value := range recordY {
		if value {
			for x := range matrix {
				matrix[x][index] = 0
			}
		}
	}
}
