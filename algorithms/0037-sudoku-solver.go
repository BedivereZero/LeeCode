package algorithms

func solveSudoku(board [][]byte) {
	row, col, blk := exist(board)
	solve(board, row, col, blk, 0)
}

func solve(board [][]byte, row [][]bool, col [][]bool, blk [][]bool, n int) bool {
	n = next(board, n)
	if n == 81 {
		return true
	}
	for i := 0; i < 9; i++ {
		if row[n/9][i] || col[n%9][i] || blk[n/27*3+n%9/3][i] {
			continue
		}
		board[n/9][n%9] = byte('1' + i)
		row[n/9][i], col[n%9][i], blk[n/27*3+n%9/3][i] = true, true, true
		if ok := solve(board, row, col, blk, n+1); ok {
			return true
		} else {
			row[n/9][i], col[n%9][i], blk[n/27*3+n%9/3][i] = false, false, false
		}
		board[n/9][n%9] = '.'
	}
	return false
}

func next(board [][]byte, i int) int {
	for i < 81 {
		if board[i/9][i%9] == '.' {
			break
		}
		i++
	}
	return i
}

func exist(board [][]byte) ([][]bool, [][]bool, [][]bool) {
	row := genMatrix(9, 9)
	col := genMatrix(9, 9)
	blk := genMatrix(9, 9)

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// skip '.'
			if board[x][y] == '.' {
				continue
			}
			// mark as row
			row[x][int(board[x][y]-'1')] = true
			// mark as column
			col[y][int(board[x][y]-'1')] = true
			// makr as block
			blk[x/3*3+y/3][int(board[x][y]-'1')] = true
		}
	}
	return row, col, blk
}

func genMatrix(x, y int) [][]bool {
	m := make([][]bool, x)
	for i := 0; i < x; i++ {
		m[i] = make([]bool, y)
	}
	return m
}
