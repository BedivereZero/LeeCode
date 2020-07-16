package algorithms

func getRow(rowIndex int) []int {
	line := make([]int, rowIndex+1)
	for x := 0; x < rowIndex+1; x++ {
		p := 1
		for y := 0; y < x+1; y++ {
			if y == 0 || y == x {
				line[y] = 1
			} else {
				p, line[y] = line[y], line[y]+p
			}
		}
	}
	return line
}
