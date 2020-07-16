package algorithms

func generate(numRows int) [][]int {
	t := make([][]int, numRows)
	for x := range t {
		t[x] = make([]int, x+1)
		for y := range t[x] {
			if y == 0 || x == y {
				t[x][y] = 1
				continue
			} else {
				t[x][y] = t[x-1][y-1] + t[x-1][y]
			}
		}
	}
	return t
}
