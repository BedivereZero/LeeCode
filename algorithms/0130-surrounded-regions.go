package algorithms

func solve(board [][]byte) {

	if len(board) == 0 {
		return
	}

	// create 2D matrix to record 'O' which connects to edge
	record := make([][]bool, len(board))
	for x := range record {
		record[x] = make([]bool, len(board[x]))
	}

	// mark all 'O' which connects to edge
	for i := 0; i < len(board); i++ {
		mark(board, record, i, 0)
		mark(board, record, i, len(board[0])-1)
	}
	for i := 0; i < len(board[0]); i++ {
		mark(board, record, 0, i)
		mark(board, record, len(board)-1, i)
	}

	// change unmarked 'O' to 'X'
	for x := range board {
		for y := range board[x] {
			if !record[x][y] {
				board[x][y] = 'X'
			}
		}
	}
}

func mark(board [][]byte, record [][]bool, x int, y int) {
	// out of bounds
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[x])-1 {
		return
	}

	// not 'X'
	if board[x][y] == 'X' {
		return
	}

	// already marked
	if record[x][y] {
		return
	}

	record[x][y] = true

	// right, up, left, down
	mark(board, record, x+1, y)
	mark(board, record, x, y+1)
	mark(board, record, x-1, y)
	mark(board, record, x, y-1)
}
