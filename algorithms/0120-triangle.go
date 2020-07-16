package algorithms

func minimumTotal(triangle [][]int) int {
	for x := 1; x < len(triangle); x++ {
		triangle[x][0] += triangle[x-1][0]
		for y := 1; y < x; y++ {
			if triangle[x-1][y-1] < triangle[x-1][y] {
				triangle[x][y] += triangle[x-1][y-1]
			} else {
				triangle[x][y] += triangle[x-1][y]
			}
		}
		triangle[x][x] += triangle[x-1][x-1]
	}
	minimum := triangle[len(triangle)-1][0]
	for y := 1; y < len(triangle); y++ {
		if minimum > triangle[len(triangle)-1][y] {
			minimum = triangle[len(triangle)-1][y]
		}
	}
	return minimum
}
