package algorithms

func solveSudoku(board [][]byte) {
	solve(board, 0)
}

func solve(board [][]byte, n int) bool {
	n = next(board, n)
	if n == 81 {
		return true
	}
	for i, f := range getCandiate(board, n) {
		if f {
			continue
		}
		board[n/9][n%9] = byte('1' + i)
		if ok := solve(board, n+1); ok {
			return true
		} else {
			board[n/9][n%9] = '.'
		}
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

func getCandiate(board [][]byte, n int) []bool {
	flag := make([]bool, 9)
	// check row
	for i := 0; i < 9; i++ {
		if board[n/9][i] != '.' {
			flag[int(board[n/9][i]-'1')] = true
		}
	}
	// check column
	for i := 0; i < 9; i++ {
		if board[i][n%9] != '.' {
			flag[int(board[i][n%9]-'1')] = true
		}
	}
	// check block
	for i := 0; i < 9; i++ {
		if board[n/27*3+i/3][n%9/3*3+i%3] != '.' {
			flag[int(board[n/27*3+i/3][n%9/3*3+i%3]-'1')] = true
		}
	}
	return flag
}
