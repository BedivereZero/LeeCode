package algorithms

func rotate(matrix [][]int) {
	// matrix: n * n
	n := len(matrix)
	// edge of part
	edgeX, edgeY := n/2, (n+1)/2
	for x := 0; x < edgeX; x++ {
		for y := 0; y < edgeY; y++ {
			matrix[x][y], matrix[y][n-1-x], matrix[n-1-x][n-1-y], matrix[n-1-y][x] = matrix[n-1-y][x], matrix[x][y], matrix[y][n-1-x], matrix[n-1-x][n-1-y]
		}
	}
}
