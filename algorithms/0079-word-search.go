package algorithms

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	footprint := createFootpoint(board)
	for x := range board {
		for y := range board[x] {
			footprint[x][y] = true
			if search(board, footprint, word, x, y) {
				return true
			} else {
				footprint[x][y] = false
			}
		}
	}
	return false
}

func createFootpoint(board [][]byte) [][]bool {
	footprint := make([][]bool, len(board))
	for index := range board {
		footprint[index] = make([]bool, len(board[index]))
	}
	return footprint
}

func search(board [][]byte, footprint [][]bool, word string, x int, y int) bool {
	if len(word) == 1 {
		return word[0] == board[x][y]
	}

	// Down
	if x + 1 < len(board) && ! footprint[x + 1][y] {
		footprint[x + 1][y] = true
		if search(board, footprint, word[1:], x + 1, y) {
			return true
		} else {
			footprint[x + 1][y] = false
		}
	}

	// Right
	if y + 1 < len(board[x]) && ! footprint[x][y + 1] {
		footprint[x][y + 1] = true
		if search(board, footprint, word[1:], x, y + 1) {
			return true
		} else {
			footprint[x][y + 1] = false
		}
	}

	// Up
	if x > 0 && ! footprint[x - 1][y] {
		footprint[x - 1][y] = true
		if search(board, footprint, word[1:], x - 1, y) {
			return true
		} else {
			footprint[x - 1][y] = false
		}
	}

	// Left
	if y > 0 && ! footprint[x][y - 1] {
		footprint[x][y - 1] = true
		if search(board, footprint, word[1:], x, y - 1) {
			return true
		} else {
			footprint[x][y - 1] = false
		}
	}

	return false
}
