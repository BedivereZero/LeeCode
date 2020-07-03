package algorithms

func solveSudoku(board [][]byte) {
	solve(board, 0)
}

func solve(board [][]byte, i int) bool {
	var x, y int
	i = next(board, i)
	if i < 81 {
		x, y = i/9, i%9
	} else {
		return true
	}
	candidate := getCandiate(board, x, y)
	for _, c := range candidate {
		board[x][y] = c
		if ok := solve(board, i+1); ok {
			return true
		} else {
			board[x][y] = '.'
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

func getCandiate(board [][]byte, x int, y int) []byte {
	flag := make([]bool, 9)
	// check row
	for i := 0; i < 9; i++ {
		if board[x][i] != '.' {
			flag[int(board[x][i]-'1')] = true
		}
	}
	// check column
	for i := 0; i < 9; i++ {
		if board[i][y] != '.' {
			flag[int(board[i][y]-'1')] = true
		}
	}
	// check block
	for i := 0; i < 9; i++ {
		if board[x/3*3+i/3][y/3*3+i%3] != '.' {
			flag[int(board[x/3*3+i/3][y/3*3+i%3]-'1')] = true
		}
	}
	var c []byte
	for i, f := range flag {
		if !f {
			c = append(c, byte('1'+i))
		}
	}
	return c
}
