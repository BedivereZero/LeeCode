package algorithms

func numIslands(grid [][]byte) int {
	num := 0
	for x, y, ok := findIsland(grid); ok; x, y, ok = findIsland(grid) {
		markIsland(x, y, grid)
		num++
	}
	return num
}

func findIsland(grid [][]byte) (int, int, bool) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '1' {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

func markIsland(x int, y int, grid [][]byte) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	markIsland(x+1, y, grid)
	markIsland(x, y+1, grid)
	markIsland(x-1, y, grid)
	markIsland(x, y-1, grid)
}
