func isValidSudoku(board [][]byte) bool {
	// valid row
	for x := 0; x < 9; x++ {
		exist := make([]bool, 9)
		for y := 0; y < 9; y++ {
			if board[x][y] != '.' {
				if exist[int(board[x][y]-'0')-1] {
					return false
				} else {
					exist[int(board[x][y]-'0')-1] = true
				}
			}
		}
	}

	// valid column
	for y := 0; y < 9; y++ {
		exist := make([]bool, 9)
		for x := 0; x < 9; x++ {
			if board[x][y] != '.' {
				if exist[int(board[x][y]-'0')-1] {
					return false
				} else {
					exist[int(board[x][y]-'0')-1] = true
				}
			}
		}
	}

	// valid block
	for i := 0; i < 9; i++ {
		exist := make([]bool, 9)
		for x := i / 3 * 3; x < i/3*3+3; x++ {
			for y := i % 3 * 3; y < i%3*3+3; y++ {
				if board[x][y] != '.' {
					if exist[int(board[x][y]-'0')-1] {
						return false
					} else {
						exist[int(board[x][y]-'0')-1] = true
					}
				}
			}
		}
	}
	return true
}
