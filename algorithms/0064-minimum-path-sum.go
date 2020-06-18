func minPathSum(grid [][]int) int {
	for x := range grid {
		for y := range grid[x] {
			if x == 0 && y == 0 {
				continue
			}
			if x == 0 && y > 0 {
				grid[x][y] += grid[x][y - 1]
				continue
			}
			if x > 0 && y == 0 {
				grid[x][y] += grid[x - 1][y]
				continue
			}
			if grid[x - 1][y] < grid[x][y - 1] {
				grid[x][y] += grid[x - 1][y]
			} else {
				grid[x][y] += grid[x][y - 1]
			}
		}
	}
	return grid[len(grid) - 1][len(grid[0]) - 1]
}
