package algorithms

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for x := 0; x < len(obstacleGrid); x++ {
		for y := 0; y < len(obstacleGrid[0]); y++ {
			if obstacleGrid[x][y] == 1 {
				obstacleGrid[x][y] = 0
				continue
			}
			if x == 0 && y == 0 {
				obstacleGrid[x][y] = 1
			}
			if x > 0 {
				obstacleGrid[x][y] += obstacleGrid[x-1][y]
			}
			if y > 0 {
				obstacleGrid[x][y] += obstacleGrid[x][y-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
