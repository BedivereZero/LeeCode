package algorithms

func calculateMinimumHP(dungeon [][]int) int {
	xLimit := len(dungeon)
	yLimit := len(dungeon[0])
	dungeon[xLimit-1][yLimit-1] = 1 - dungeon[xLimit-1][yLimit-1]
	if dungeon[xLimit-1][yLimit-1] < 1 {
		dungeon[xLimit-1][yLimit-1] = 1
	}
	for x := xLimit - 2; x >= 0; x-- {
		dungeon[x][yLimit-1] = dungeon[x+1][yLimit-1] - dungeon[x][yLimit-1]
		if dungeon[x][yLimit-1] < 1 {
			dungeon[x][yLimit-1] = 1
		}
	}
	for y := yLimit - 2; y >= 0; y-- {
		dungeon[xLimit-1][y] = dungeon[xLimit-1][y+1] - dungeon[xLimit-1][y]
		if dungeon[xLimit-1][y] < 1 {
			dungeon[xLimit-1][y] = 1
		}
	}
	for x := xLimit - 2; x >= 0; x-- {
		for y := yLimit - 2; y >= 0; y-- {
			if dungeon[x+1][y] < dungeon[x][y+1] {
				dungeon[x][y] = dungeon[x+1][y] - dungeon[x][y]
			} else {
				dungeon[x][y] = dungeon[x][y+1] - dungeon[x][y]
			}
			if dungeon[x][y] < 1 {
				dungeon[x][y] = 1
			}
		}
	}
	return dungeon[0][0]
}
